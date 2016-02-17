package gojieba

import "C"
import "unsafe"

func cstrings(x **C.char) []string {
	var s []string
	eltSize := unsafe.Sizeof(*x)
	for *x != nil {
		s = append(s, C.GoString(*x))
		x = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + eltSize))
	}
	return s
}
