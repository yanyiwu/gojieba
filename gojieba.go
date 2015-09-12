package gojieba

/*
#cgo CXXFLAGS: -DLOGGER_LEVEL=LL_WARN -O3 -Wall
#include "cjieba.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type GoJieba struct {
	jieba C.CJieba
}

func New(dict_path, hmm_path, user_dict_path string) GoJieba {
	var x GoJieba
	x.jieba = C.NewCJieba(C.CString(dict_path), C.CString(hmm_path), C.CString(user_dict_path))
	return x
}

func (x GoJieba) Free() {
	C.FreeCJieba(x.jieba)
}

func (x GoJieba) Cut(s string) string {
	var res *C.char = C.Cut(x.jieba, C.CString(s), C.CString("/"))
	gs := C.GoString(res)
	C.free(unsafe.Pointer(res))
	return gs
}
