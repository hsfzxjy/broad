package broad

import (
	"sync/atomic"
	"unsafe"
)

type _GetResult int

const (
	ngrOK _GetResult = iota
	ngrStarved
	ngrFull
)

type node[T any] interface {
	LoadNext() node[T]
	StoreNext(unsafe.Pointer)
	AsRaw() unsafe.Pointer
	Len() uintptr
	Get(idx uintptr) (*T, _GetResult)
	Append(elt T)
	Extend(elts []T)
	Avail() uintptr
}

const (
	tagInline = 0b00
	tagSlice  = 0b01
	tagBlock  = 0b11
	tagMask   = 0b11
)

var _sentinel uintptr
var (
	pInlineNil = unsafe.Pointer(&_sentinel)
	pSliceNil  = unsafe.Add(unsafe.Pointer(&_sentinel), tagSlice)
	pBlockNil  = unsafe.Add(unsafe.Pointer(&_sentinel), tagBlock)
)

func parseRawNode[T any](p unsafe.Pointer) node[T] {
	nodeAddr := (*uintptr)(unsafe.Pointer(uintptr(p) &^ tagMask))
	tag := atomic.LoadUintptr(nodeAddr) & tagMask
	switch tag {
	case tagInline:
		return (*inlineNode[T])(unsafe.Pointer(nodeAddr))
	case tagBlock:
		return (*blockNode[T])(unsafe.Pointer(nodeAddr))
	case tagSlice:
		return (*sliceNode[T])(unsafe.Pointer(nodeAddr))
	default:
		return nil
	}
}

type blockNode[T any] struct {
	next   unsafe.Pointer
	lenCap uintptr
	data   unsafe.Pointer
}

const wordBits = unsafe.Sizeof(uintptr(0)) * 8
const halfWordBits = wordBits / 2

func newBlockNode[T any](data []T) *blockNode[T] {
	lenCap := uintptr(len(data))<<halfWordBits | uintptr(cap(data))
	return &blockNode[T]{
		data:   unsafe.Pointer(unsafe.SliceData(data)),
		lenCap: lenCap,
		next:   pBlockNil,
	}
}

func (n *blockNode[T]) LoadNext() node[T] {
	ptr := atomic.LoadPointer(&n.next)
	if ptr == pBlockNil {
		return nil
	}
	return parseRawNode[T](ptr)
}

func (n *blockNode[T]) StoreNext(ptr unsafe.Pointer) {
	ptr = unsafe.Add(ptr, tagBlock)
	atomic.StorePointer(&n.next, ptr)
}

func (n *blockNode[T]) AsRaw() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n *blockNode[T]) Len() uintptr {
	lenCap := atomic.LoadUintptr(&n.lenCap)
	return lenCap >> halfWordBits
}

func (n *blockNode[T]) Get(idx uintptr) (*T, _GetResult) {
	lenCap := atomic.LoadUintptr(&n.lenCap)
	len := lenCap >> halfWordBits
	cap := lenCap & (1<<halfWordBits - 1)
	switch {
	case idx < len:
		var zero T
		return (*T)(unsafe.Add(n.data, uintptr(idx)*unsafe.Sizeof(zero))), ngrOK
	case idx < cap:
		return nil, ngrStarved
	case len == cap:
		return nil, ngrFull
	default:
		panic("unreachable: index out of range")
	}
}

func (n *blockNode[T]) Append(elt T) {
	lenCap := atomic.LoadUintptr(&n.lenCap)
	l := lenCap >> halfWordBits
	c := lenCap & (1<<halfWordBits - 1)
	if l >= c {
		panic("unreachable: no space to append")
	}
	var zero T
	ptr := (*T)(unsafe.Add(
		unsafe.Pointer(n.data),
		uintptr(l)*unsafe.Sizeof(zero),
	))
	*ptr = elt
	newLenCap := (l+1)<<halfWordBits | c
	if !atomic.CompareAndSwapUintptr(&n.lenCap, lenCap, newLenCap) {
		panic("concurrent Append() detected")
	}
}

func (n *blockNode[T]) Extend(elts []T) {
	lenCap := atomic.LoadUintptr(&n.lenCap)
	l := int(lenCap >> halfWordBits)
	c := int(lenCap & (1<<halfWordBits - 1))
	newL := len(elts) + l
	newLenCap := uintptr(newL)<<halfWordBits | uintptr(c)
	var zero T
	slice := unsafe.Slice(
		(*T)(unsafe.Add(
			unsafe.Pointer(n.data),
			uintptr(l)*unsafe.Sizeof(zero),
		)),
		len(elts),
	)
	copy(slice, elts)
	if !atomic.CompareAndSwapUintptr(&n.lenCap, lenCap, newLenCap) {
		panic("concurrent Extend() detected")
	}
}

func (n *blockNode[T]) Avail() uintptr {
	lenCap := atomic.LoadUintptr(&n.lenCap)
	len := lenCap >> halfWordBits
	cap := lenCap & (1<<halfWordBits - 1)
	return cap - len
}

type sliceNode[T any] struct {
	next unsafe.Pointer
	len  uintptr
	data unsafe.Pointer
}

func newSliceNode[T any](data []T) *sliceNode[T] {
	len := len(data)
	return &sliceNode[T]{
		data: unsafe.Pointer(unsafe.SliceData(data)),
		len:  uintptr(len),
		next: pSliceNil,
	}
}

func (n *sliceNode[T]) LoadNext() node[T] {
	ptr := atomic.LoadPointer(&n.next)
	if ptr == pSliceNil {
		return nil
	}
	return parseRawNode[T](ptr)
}

func (n *sliceNode[T]) StoreNext(ptr unsafe.Pointer) {
	ptr = unsafe.Add(ptr, tagSlice)
	atomic.StorePointer(&n.next, ptr)
}

func (n *sliceNode[T]) AsRaw() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n *sliceNode[T]) Len() uintptr {
	return n.len
}

func (n *sliceNode[T]) Get(idx uintptr) (*T, _GetResult) {
	if idx >= n.len {
		return nil, ngrFull
	}
	var zero T
	return (*T)(unsafe.Add(n.data, uintptr(idx)*unsafe.Sizeof(zero))), ngrOK
}

func (n *sliceNode[T]) Append(elt T) {
	panic("unreachable: sliceNode.Append not implemented")
}

func (n *sliceNode[T]) Extend(elts []T) {
	panic("unreachable: sliceNode.Extend not implemented")
}

func (n *sliceNode[T]) Avail() uintptr {
	return 0
}

type inlineNode[T any] struct {
	next unsafe.Pointer
	data T
}

func newInlineNode[T any](data T) *inlineNode[T] {
	return &inlineNode[T]{data: data, next: pInlineNil}
}

func (n *inlineNode[T]) LoadNext() node[T] {
	ptr := atomic.LoadPointer(&n.next)
	if ptr == pInlineNil {
		return nil
	}
	return parseRawNode[T](ptr)
}

func (n *inlineNode[T]) StoreNext(ptr unsafe.Pointer) {
	ptr = unsafe.Add(ptr, tagInline)
	atomic.StorePointer(&n.next, ptr)
}

func (n *inlineNode[T]) AsRaw() unsafe.Pointer {
	return unsafe.Pointer(n)
}

func (n *inlineNode[T]) Get(idx uintptr) (*T, _GetResult) {
	if idx > 0 {
		return nil, ngrFull
	}
	return &n.data, ngrOK
}

func (n *inlineNode[T]) Extend(elts []T) {
	panic("unreachable: inlineNode.Extend not implemented")
}

func (n *inlineNode[T]) Append(elt T) {
	panic("unreachable: inlineNode.Append not implemented")
}

func (n *inlineNode[T]) Avail() uintptr {
	return 0
}

func (n *inlineNode[T]) Len() uintptr {
	return 1
}
