package gojieba

/*
#cgo CXXFLAGS: -DLOGGER_LEVEL=LL_WARN -O3 -Wall
#include <stdlib.h>
#include "jieba.h"
*/
import "C"
import "unsafe"

type GoJieba struct {
	jieba C.Jieba
}

func New(dict_path, hmm_path, user_dict_path string) GoJieba {
	var x GoJieba
	x.jieba = C.NewJieba(C.CString(dict_path), C.CString(hmm_path), C.CString(user_dict_path))
	return x
}

func (x GoJieba) Free() {
	C.FreeJieba(x.jieba)
}

func (x GoJieba) Cut(s string) []string {
	var a **C.char = C.Cut(x.jieba, C.CString(s))
	res := CStrings(a)
	C.FreeWords(a)
	return res
}

func sptr(p uintptr) *C.char {
	return *(**C.char)(unsafe.Pointer(p))
}

func CStrings(x **C.char) []string {
	var s []string
	for p := uintptr(unsafe.Pointer(x)); sptr(p) != nil; p += unsafe.Sizeof(uintptr(0)) {
		s = append(s, C.GoString(sptr(p)))
	}
	return s
}
