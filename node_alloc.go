package broad

import "unsafe"

//go:generate go run node_mksizes.go

func allocNodeInit[T any]() node[T] {
	var dummy T
	var sizeT = unsafe.Sizeof(dummy)
	if sizeT == 0 {
		return newZSTNode[T](1)
	}
	return newBlockNode([]T(nil))
}

func allocNodeForElt[T any](elt T) node[T] {
	var sizeT = unsafe.Sizeof(elt)
	if sizeT == 0 {
		panic("unreachable: zero-sized type")
	}
	var cap_ = byte(1)
	if sizeT <= _NSizes {
		cap_ = sizes[sizeT].Cap
	}
	if cap_ == 1 {
		return newInlineNode(elt)
	}
	data := make([]T, 1, cap_)
	data[0] = elt
	return newBlockNode(data)
}

func allocNodeForSlice[T any](elts []T) node[T] {
	var sizeT = unsafe.Sizeof(elts[0])
	if sizeT == 0 {
		panic("unreachable: zero-sized type")
	}
	return newSliceNode(elts)
}
