package lq

import (
	"iter"
	"sync/atomic"
	"unsafe"
)

type node[T any] interface {
	LoadNext() node[T]
	StoreNext(unsafe.Pointer)
	AsRaw() unsafe.Pointer
	Size() int
	Get(idx int) (T, bool)
}

var _sentinel uintptr
var (
	pSingleNil = unsafe.Pointer(&_sentinel)
	pSliceNil  = unsafe.Add(unsafe.Pointer(&_sentinel), 1)
)

func castRawNode[T any](p unsafe.Pointer) node[T] {
	nodeAddr := (*uintptr)(unsafe.Pointer(uintptr(p) &^ 0x1))
	tag := atomic.LoadUintptr(nodeAddr)
	if tag&1 == 0 {
		return (*singleNode[T])(unsafe.Pointer(nodeAddr))
	} else {
		return (*sliceNode[T])(unsafe.Pointer(nodeAddr))
	}

}

type singleNode[T any] struct {
	next unsafe.Pointer
	v    T
}

func newSingleNode[T any](v T) *singleNode[T] {
	return &singleNode[T]{v: v, next: pSingleNil}
}

func (n *singleNode[T]) LoadNext() node[T] {
	ptr := atomic.LoadPointer(&n.next)
	if ptr == pSingleNil {
		return nil
	}
	return castRawNode[T](ptr)
}

func (n *singleNode[T]) StoreNext(ptr unsafe.Pointer) {
	atomic.StorePointer(&n.next, ptr)
}

func (n *singleNode[T]) AsRaw() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n *singleNode[T]) Size() int {
	return 1
}

func (n *singleNode[T]) Get(idx int) (T, bool) {
	if idx == 0 {
		return n.v, true
	}
	var zero T
	return zero, false
}

type sliceNode[T any] struct {
	next unsafe.Pointer
	size int
	data *T
}

func newSliceNode[T any](data []T) *sliceNode[T] {
	return &sliceNode[T]{data: unsafe.SliceData(data), size: len(data), next: pSliceNil}
}

func (n *sliceNode[T]) LoadNext() node[T] {
	ptr := atomic.LoadPointer(&n.next)
	if ptr == pSliceNil {
		return nil
	}
	return castRawNode[T](ptr)
}

func (n *sliceNode[T]) StoreNext(ptr unsafe.Pointer) {
	ptr = unsafe.Add(ptr, 1)
	atomic.StorePointer(&n.next, ptr)
}

func (n *sliceNode[T]) AsRaw() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n *sliceNode[T]) Size() int {
	return n.size
}

func (n *sliceNode[T]) Get(idx int) (T, bool) {
	var zero T
	if 0 <= idx && idx < n.size {
		return *(*T)(unsafe.Add(unsafe.Pointer(n.data), uintptr(idx)*unsafe.Sizeof(zero))), true
	}
	return zero, false
}

// [Queue] is a single-producer, multi-consumer queue.
//
// Use [Queue.Push] or [Queue.PushSlice] to push items into the queue.
// Pushes are non-blocking and always succeed unless the queue is closed.
// However, concurrent pushing is not safe.
//
// Use [Queue.Listen] to get a listener, which yields items in the order they were pushed.
// It's safe to have multiple listeners, and have them listening concurrently.
// Each listener will receive all items pushed after it was created.
//
// There's no need to explicitly unregister listeners, so simply discard them when out of use.
//
// Use [Queue.Close] to close the queue. After closing, no more items can be pushed into the queue.
type Queue[T any] struct {
	notifier unsafe.Pointer

	tail node[T]

	closed atomic.Bool
}

// New creates a new [Queue].
func New[T any]() *Queue[T] {
	q := &Queue[T]{}
	ch := make(chan struct{})
	q.notifier = *(*unsafe.Pointer)(unsafe.Pointer(&ch))
	q.tail = newSliceNode[T](nil)
	return q
}

// Push pushes an item into the queue.
// Returns false if the queue is closed.
func (q *Queue[T]) Push(v T) bool {
	if q.closed.Load() {
		return false
	}
	newNode := newSingleNode(v)
	oldTail := q.tail
	oldTail.StoreNext(newNode.AsRaw())
	q.tail = newNode
	q.updateNotifier()
	return true
}

// PushSlice pushes a slice of items into the queue.
// Returns false if the queue is closed.
//
// The underlying array is not copied, so the caller should not modify the array after pushing it into the queue.
func (q *Queue[T]) PushSlice(s []T) bool {
	if q.closed.Load() {
		return false
	}
	if len(s) != 0 {
		newNode := newSliceNode(s)

		oldTail := q.tail
		oldTail.StoreNext(newNode.AsRaw())
		q.tail = newNode
	}
	q.updateNotifier()
	return true
}

var closedNotifier = make(chan struct{})
var closedNotifierPtr = *(*unsafe.Pointer)(unsafe.Pointer(&closedNotifier))

func init() { close(closedNotifier) }

// Close closes the queue.
// After closing, no more items can be pushed into the queue.
func (q *Queue[T]) Close() {
	if q.closed.CompareAndSwap(false, true) {
		oldp := atomic.SwapPointer(&q.notifier, closedNotifierPtr)
		oldp = unsafe.Pointer(uintptr(oldp) &^ 0x1)
		oldNotif := *(*chan struct{})(unsafe.Pointer(&oldp))
		if oldNotif != closedNotifier {
			close(oldNotif)
		}
	}

}

// Listen creates a listener for the queue. The listener will receive all items pushed after it was created.
func (q *Queue[T]) Listen() *Listener[T] {
	tail := q.tail
	return &Listener[T]{
		q:      q,
		head:   tail,
		offset: tail.Size(),
	}
}

func (q *Queue[T]) loadNotifier() chan struct{} {
	cur := atomic.LoadPointer(&q.notifier)
	if uintptr(cur)&1 == 1 {
		cur = unsafe.Pointer(uintptr(cur) &^ 0x1)
	} else {
		atomic.CompareAndSwapPointer(&q.notifier, cur, unsafe.Add(cur, 1))
	}
	cur = unsafe.Pointer(uintptr(cur) &^ 0x1)
	notif := *(*chan struct{})(unsafe.Pointer(&cur))
	return notif
}

func (q *Queue[T]) updateNotifier() {
	oldPtr := atomic.LoadPointer(&q.notifier)
	if uintptr(oldPtr)&1 == 1 {
		chptr := unsafe.Pointer(uintptr(oldPtr) &^ 0x1)
		notif := *(*chan struct{})(unsafe.Pointer(&chptr))
		if notif != closedNotifier {
			close(notif)
		} else {
			return
		}
		ch := make(chan struct{})
		if atomic.CompareAndSwapPointer(&q.notifier, oldPtr, *(*unsafe.Pointer)(unsafe.Pointer(&ch))) {
			return
		} else {
			panic("concurrent Push*() or Close() detected")
		}
	}
}

// Listener receives items from a [Queue].
type Listener[T any] struct {
	q    *Queue[T]
	head node[T]

	offset int
}

// Next returns the next item in the queue.
// Returns false if there's no more item in the queue (queue closed).
// It's not safe to call Next concurrently with other Next and Iter calls.
func (o *Listener[T]) Next() (T, bool) {
	item, ok := o.head.Get(o.offset)
	if ok {
		o.offset++
		return item, true
	}
	var next node[T]
LOAD_NEXT:
	next = o.head.LoadNext()
SWITCH_NEXT:
	if next == nil {
		notif := o.q.loadNotifier()
		next = o.head.LoadNext()
		if next != nil {
			goto SWITCH_NEXT
		}
		if notif == closedNotifier {
			var zero T
			return zero, false
		}
		<-notif
		goto LOAD_NEXT
	}

	item, _ = next.Get(0)
	o.head = next
	o.offset = 1
	return item, true
}

// Iter returns an iterator that yields items from the queue.
// It's not safe to call Iter concurrently with other Next and Iter calls.
func (o *Listener[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			item, ok := o.Next()
			if !ok {
				break
			}
			if !yield(item) {
				break
			}
		}
	}
}
