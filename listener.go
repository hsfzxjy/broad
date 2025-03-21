package broad

import (
	"context"
	"iter"
)

// A [Listener] receives items in order from a [Caster] that are pushed after the moment the [Listener] was created.
//
// [Listener] provides [Listener.Next], [Listener.NextCtx], [Listener.Iter], [Listener.IterCtx] and [Listener.IterUntilBlocked] methods to yield items, depending on use cases.
//
// It's NOT SAFE to call [Listener.Next]* or [Listener.Iter]* concurrently on the same [Listener] instance.
// However, concurrent use of multiple [Listener] instances is safe.
type Listener[T any] struct {
	c    *Caster[T]
	head node[T]

	offset uintptr
}

// NextCtx receives the next item from the associated [Caster], with ctx as a canceler.
// ok being false indicates either of the following situations:
//
// - The [Caster] is closed, and no more unconsumed items are available; or
//
// - ctx is canceled.
//
// Otherwise, ok is true, and val is the next item in the [Caster].
//
// The method will block until either a new item is available or ctx is canceled.
//
// NOTE: Passing an already canceled ctx won't always result in ok == false, as
// the method only checks the context when being blocked.
func (l *Listener[T]) NextCtx(ctx context.Context) (val T, ok bool) {
	ptr, ok := l.next(ctx.Done())
	if !ok {
		return val, false
	}
	return *ptr, true
}

// NextCtx receives the next item from the associated [Caster].
// ok is false if and only if the [Caster] is closed,
// and no more unconsumed items are available.
//
// Otherwise, ok is true, and val is the next item in the [Caster].
//
// The method will block until a new item is available.
func (l *Listener[T]) Next() (val T, ok bool) {
	ptr, ok := l.next(nil)
	if !ok {
		return val, false
	}
	return *ptr, true
}

func (l *Listener[T]) next(doneCh <-chan struct{}) (*T, bool) {
LOAD_AGAIN:
	item, res := l.head.Get(l.offset)
	if res == ngrOK {
		l.offset++
		return item, true
	}
	if res == ngrStarved {
		notif := l.c.loadNotifier()
		item, res := l.head.Get(l.offset)
		if res == ngrOK {
			l.offset++
			return item, true
		}
		if notif == closedNotifier {
			return nil, false
		}
		if doneCh != nil {
			select {
			case <-notif:
				goto LOAD_AGAIN
			case <-doneCh:
				return nil, false
			}
		} else {
			<-notif
			goto LOAD_AGAIN
		}
	}
	var next node[T]
LOAD_NEXT:
	next = l.head.LoadNext()
SWITCH_NEXT:
	if next == nil {
		notif := l.c.loadNotifier()
		next = l.head.LoadNext()
		if next != nil {
			goto SWITCH_NEXT
		}
		if notif == closedNotifier {
			return nil, false
		}
		if doneCh != nil {
			select {
			case <-notif:
			case <-doneCh:
				return nil, false
			}
		} else {
			<-notif
		}
		goto LOAD_NEXT
	}

	item, _ = next.Get(0)
	l.head = next
	l.offset = 1
	return item, true
}

// IterCtx returns an iter.Seq[T] that yields items from the [Caster].
//
// The returned iterator can be interrupted by canceling the provided context.
func (l *Listener[T]) IterCtx(ctx context.Context) iter.Seq[T] {
	return func(yield func(T) bool) {
		doneCh := ctx.Done()
		for {
			item, ok := l.next(doneCh)
			if !ok {
				break
			}
			if !yield(*item) {
				break
			}
		}
	}
}

// Iter returns an iter.Seq[T] that yields items from the [Caster].
//
// The returned iterator runs until the [Caster] is closed and all items are consumed.
func (l *Listener[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			item, ok := l.next(nil)
			if !ok {
				break
			}
			if !yield(*item) {
				break
			}
		}
	}
}

// IterUntilBlocked returns an iter.Seq[T] that yields items from the [Caster].
//
// The returned iterator runs until no unconsumed items left.
func (l *Listener[T]) IterUntilBlocked() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			item, ok := l.next(closedNotifier)
			if !ok {
				break
			}
			if !yield(*item) {
				break
			}
		}
	}
}
