package gojieba

/*
#include "jieba.h"
*/
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

func convertWords(s string, words *C.Word) []Word {
	result := make([]Word, 0)
	x := words
	eltSize := unsafe.Sizeof(*x)
	start := 0
	end := 0
	for (*x).len != 0 {
		start = int((*x).offset)
		end = start + int((*x).len)
		result = append(result, Word{s[start:end], start, end})
		x = (*C.Word)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + eltSize))
	}
	return result
}
