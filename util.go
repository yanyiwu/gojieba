package gojieba

import "C"
import "unsafe"

func sptr(p uintptr) *C.char {
	return *(**C.char)(unsafe.Pointer(p))
}

func cstrings(x **C.char) []string {
	var s []string
	for p := uintptr(unsafe.Pointer(x)); sptr(p) != nil; p += unsafe.Sizeof(uintptr(0)) {
		s = append(s, C.GoString(sptr(p)))
	}
	return s
}
