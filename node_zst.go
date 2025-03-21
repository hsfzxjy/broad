package broad

import (
	"sync/atomic"
	"unsafe"
)

type zstNode[T any] struct {
	data T
	size atomic.Uintptr
}

func newZSTNode[T any](initSize uintptr) *zstNode[T] {
	n := &zstNode[T]{}
	n.size.Store(initSize)
	return n
}

func (n *zstNode[T]) LoadNext() node[T]            { panic("not implemented") }
func (n *zstNode[T]) StoreNext(ptr unsafe.Pointer) { panic("not implemented") }

func (n *zstNode[T]) AsRaw() unsafe.Pointer { return unsafe.Pointer(n) }
func (n *zstNode[T]) Get(idx uintptr) (*T, _GetResult) {
	size := n.size.Load()
	if uintptr(idx) < size {
		return &n.data, ngrOK
	}
	return nil, ngrStarved
}

func (n *zstNode[T]) Append(elt T) {
	newSize := n.size.Add(1)
	if newSize == 0 {
		panic("zstNode overflow")
	}
}

func (n *zstNode[T]) Extend(elts []T) {
	newSize := n.size.Add(uintptr(len(elts)))
	if newSize < uintptr(len(elts)) {
		panic("zstNode overflow")
	}
}

func (n *zstNode[T]) Avail() uintptr {
	return ^uintptr(0) - n.size.Load()
}

func (n *zstNode[T]) Len() uintptr {
	return n.size.Load()
}
