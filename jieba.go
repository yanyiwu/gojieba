package gojieba

/*
#cgo CXXFLAGS: -I./deps -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#include <stdlib.h>
#include "jieba.h"
*/
import "C"
import "unsafe"

type TokenizeMode int

const (
	DefaultMode TokenizeMode = iota
	SearchMode
)

type Word struct {
	Str   string
	Start int
	End   int
}

type Jieba struct {
	jieba C.Jieba
}

func NewJieba(paths ...string) *Jieba {
	dictpaths := getDictPaths(paths...)
	dpath := C.CString(dictpaths[0])
	defer C.free(unsafe.Pointer(dpath))
	hpath := C.CString(dictpaths[1])
	defer C.free(unsafe.Pointer(hpath))
	upath := C.CString(dictpaths[2])
	defer C.free(unsafe.Pointer(upath))
	return &Jieba{
		C.NewJieba(dpath, hpath, upath),
	}
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

func (x *Jieba) AddWord(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.AddWord(x.jieba, cstr)
}

func (x *Jieba) Tokenize(s string, mode TokenizeMode, hmm bool) []Word {
	c_int_hmm := 0
	if hmm {
		c_int_hmm = 1
	}
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words *C.Word = C.Tokenize(x.jieba, cstr, C.TokenizeMode(mode), C.int(c_int_hmm))
	defer C.free(unsafe.Pointer(words))
	return convertWords(s, words)
}
