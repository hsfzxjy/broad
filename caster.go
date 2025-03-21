package broad

import (
	"sync/atomic"
	"unsafe"
)

// [Caster] represents a single-producer broadcaster.
//
// Use [Caster.Push] or [Caster.PushSlice] to push items into the [Caster].
// Pushes are non-blocking and always succeed unless the [Caster] is closed.
// However, concurrent pushing is not safe.
//
// Use [Caster.Listen] to create a listener.
// Each listener receives all items that are pushed after the moment the listener was created.
// It's safe to have multiple listeners, and have them listening concurrently.
//
// Listeners need no explicit unregisteration, so simply discard them at end of use.
//
// Use [Caster.Close] to close the [Caster].
// After being closed, no more items can be pushed into the [Caster].
type Caster[T any] struct {
	notifier unsafe.Pointer

	tail node[T]

	closed atomic.Bool
}

// New creates a new [Caster].
func New[T any]() *Caster[T] {
	c := &Caster[T]{}
	ch := make(chan struct{})
	c.notifier = *(*unsafe.Pointer)(unsafe.Pointer(&ch))
	c.tail = allocNodeInit[T]()
	return c
}

// Push pushes an item into the [Caster].
// Returns false if the [Caster] is closed.
func (l *Caster[T]) Push(v T) bool {
	if l.closed.Load() {
		return false
	}
	avail := l.tail.Avail()
	if avail > 0 {
		l.tail.Append(v)
	} else {
		newNode := allocNodeForElt(v)
		oldTail := l.tail
		oldTail.StoreNext(newNode.AsRaw())
		l.tail = newNode
	}
	l.updateNotifier()
	return true
}

// PushSlice pushes a slice of items into the [Caster].
// Returns false if the [Caster] is closed.
//
// NOTE: The underlying memory of s MIGHT be reused by the [Caster].
// Hence, caller SHOULD NOT modify s after calling PushSlice.
func (l *Caster[T]) PushSlice(s []T) bool {
	if l.closed.Load() {
		return false
	}
	if len(s) == 0 {
		return true
	}
	avail := l.tail.Avail()
	if avail >= uintptr(len(s)) {
		l.tail.Extend(s)
	} else {
		if avail > 0 {
			l.tail.Extend(s[:avail])
			s = s[avail:]
		}
		newNode := allocNodeForSlice(s)
		oldTail := l.tail
		oldTail.StoreNext(newNode.AsRaw())
		l.tail = newNode
	}

	l.updateNotifier()
	return true
}

var closedNotifier = make(chan struct{})
var closedNotifierPtr = *(*unsafe.Pointer)(unsafe.Pointer(&closedNotifier))

func init() { close(closedNotifier) }

// Close closes the [Caster].
// After being closed, no more items can be pushed into the [Caster].
func (l *Caster[T]) Close() {
	if l.closed.CompareAndSwap(false, true) {
		oldp := atomic.SwapPointer(&l.notifier, closedNotifierPtr)
		oldp = unsafe.Pointer(uintptr(oldp) &^ 0x1)
		oldNotif := *(*chan struct{})(unsafe.Pointer(&oldp))
		if oldNotif != closedNotifier {
			close(oldNotif)
		}
	}
}

// Listen creates a listener.
// The listener will receive all items pushed after the moment it was created.
func (l *Caster[T]) Listen() *Listener[T] {
	tail := l.tail
	return &Listener[T]{
		c:      l,
		head:   tail,
		offset: tail.Len(),
	}
}

func (l *Caster[T]) loadNotifier() chan struct{} {
	cur := atomic.LoadPointer(&l.notifier)
	if uintptr(cur)&1 == 1 {
		cur = unsafe.Pointer(uintptr(cur) &^ 0x1)
	} else {
		atomic.CompareAndSwapPointer(&l.notifier, cur, unsafe.Add(cur, 1))
	}
	cur = unsafe.Pointer(uintptr(cur) &^ 0x1)
	notif := *(*chan struct{})(unsafe.Pointer(&cur))
	return notif
}

func (l *Caster[T]) updateNotifier() {
	oldPtr := atomic.LoadPointer(&l.notifier)
	if uintptr(oldPtr)&1 == 1 {
		chptr := unsafe.Pointer(uintptr(oldPtr) &^ 0x1)
		notif := *(*chan struct{})(unsafe.Pointer(&chptr))
		if notif != closedNotifier {
			close(notif)
		} else {
			return
		}
		ch := make(chan struct{})
		if atomic.CompareAndSwapPointer(&l.notifier, oldPtr, *(*unsafe.Pointer)(unsafe.Pointer(&ch))) {
			return
		} else {
			panic("concurrent Push*() or Close() detected")
		}
	}
}
