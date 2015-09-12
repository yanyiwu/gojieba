package gojieba

/*
#cgo CXXFLAGS: -DLOGGER_LEVEL=LL_WARN -O3 -Wall
#include <stdlib.h>
#include "jieba.h"
*/
import "C"

type Jieba struct {
	jieba C.Jieba
}

func New(dict_path, hmm_path, user_dict_path string) Jieba {
	var x Jieba
	x.jieba = C.NewJieba(C.CString(dict_path), C.CString(hmm_path), C.CString(user_dict_path))
	return x
}

func (x Jieba) Free() {
	C.FreeJieba(x.jieba)
}

func (x Jieba) Cut(s string, hmm bool) []string {
	c_int_hmm := 0
	if hmm {
		c_int_hmm = 1
	}
	var words **C.char = C.Cut(x.jieba, C.CString(s), C.int(c_int_hmm))
	res := cstrings(words)
	C.FreeWords(words)
	return res
}
