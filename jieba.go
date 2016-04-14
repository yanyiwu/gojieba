package gojieba

/*
#cgo CXXFLAGS: -I./deps -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#include <stdlib.h>
#include "jieba.h"
*/
import "C"
import "unsafe"

type Jieba struct {
	jieba C.Jieba
}

func NewJieba(paths ...string) *Jieba {
	dict_path, hmm_path, user_dict_path := getDictPaths(paths...)
	dpath := C.CString(dict_path)
	defer C.free(unsafe.Pointer(dpath))
	hpath := C.CString(hmm_path)
	defer C.free(unsafe.Pointer(hpath))
	upath := C.CString(user_dict_path)
	defer C.free(unsafe.Pointer(upath))
	return &Jieba{
		C.NewJieba(dpath, hpath, upath),
	}
}

func getDictPaths(args ...string) (string, string, string) {
	dicts := [3]string{
		DICT_PATH,
		HMM_PATH,
		USER_DICT_PATH,
	}
	for i := 0; i < len(args) && i < len(dicts); i++ {
		dicts[i] = args[i]
	}
	return dicts[0], dicts[1], dicts[2]
}

func (x *Jieba) Free() {
	C.FreeJieba(x.jieba)
}

func (x *Jieba) Cut(s string, hmm bool) []string {
	c_int_hmm := 0
	if hmm {
		c_int_hmm = 1
	}
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.Cut(x.jieba, cstr, C.int(c_int_hmm))
	defer C.FreeWords(words)
	res := cstrings(words)
	return res
}

func (x *Jieba) CutAll(s string) []string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.CutAll(x.jieba, cstr)
	defer C.FreeWords(words)
	res := cstrings(words)
	return res
}

func (x *Jieba) CutForSearch(s string, hmm bool) []string {
	c_int_hmm := 0
	if hmm {
		c_int_hmm = 1
	}
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.CutForSearch(x.jieba, cstr, C.int(c_int_hmm))
	defer C.FreeWords(words)
	res := cstrings(words)
	return res
}

func (x *Jieba) Tag(s string) []string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.Tag(x.jieba, cstr)
	defer C.FreeWords(words)
	res := cstrings(words)
	return res
}
