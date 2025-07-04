package gojieba

/*
#cgo CXXFLAGS: -I./deps/cppjieba/include -I./deps/cppjieba/deps/limonp/include -DLOGGING_LEVEL=LL_WARNING -O3 -Wno-deprecated -Wno-unused-variable -std=c++11
#include <stdlib.h>
#include "jieba.h"
*/
import "C"
import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"unsafe"
)

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
	freed int32
}

func NewJieba(paths ...string) *Jieba {
	dictpaths := getDictPaths(paths...)

	// check if the dictionary files exist
	for _, path := range dictpaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic(fmt.Sprintf("Dictionary file does not exist: %s", path))
		}
	}

	dpath, hpath, upath, ipath, spath := C.CString(dictpaths[0]), C.CString(dictpaths[1]), C.CString(dictpaths[2]), C.CString(dictpaths[3]), C.CString(dictpaths[4])
	defer C.free(unsafe.Pointer(dpath))
	defer C.free(unsafe.Pointer(hpath))
	defer C.free(unsafe.Pointer(upath))
	defer C.free(unsafe.Pointer(ipath))
	defer C.free(unsafe.Pointer(spath))
	jieba := &Jieba{
		C.NewJieba(
			dpath,
			hpath,
			upath,
			ipath,
			spath,
		),
		0,
	}
	// set finalizer to free the memory when the object is garbage collected
	runtime.SetFinalizer(jieba, (*Jieba).Free)
	return jieba
}

func (x *Jieba) Free() {
	if atomic.CompareAndSwapInt32(&x.freed, 0, 1) { // only free once
		C.FreeJieba(x.jieba)
	}
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

func (x *Jieba) AddWordEx(s string, freq int, tag string) {
	cstr := C.CString(s)
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	defer C.free(unsafe.Pointer(cstr))
	C.AddWordEx(x.jieba, cstr, C.int(freq), ctag)
}

func (x *Jieba) RemoveWord(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.RemoveWord(x.jieba, cstr)
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

type WordWeight struct {
	Word   string
	Weight float64
}

func (x *Jieba) Extract(s string, topk int) []string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.Extract(x.jieba, cstr, C.int(topk))
	res := cstrings(words)
	defer C.FreeWords(words)
	return res
}

func (x *Jieba) ExtractWithWeight(s string, topk int) []WordWeight {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	words := C.ExtractWithWeight(x.jieba, cstr, C.int(topk))
	p := unsafe.Pointer(words)
	res := cwordweights((*C.struct_CWordWeight)(p))
	defer C.FreeWordWeights(words)
	return res
}

func cwordweights(x *C.struct_CWordWeight) []WordWeight {
	var s []WordWeight
	eltSize := unsafe.Sizeof(*x)
	for (*x).word != nil {
		ww := WordWeight{
			C.GoString(((C.struct_CWordWeight)(*x)).word),
			float64((*x).weight),
		}
		s = append(s, ww)
		x = (*C.struct_CWordWeight)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + eltSize))
	}
	return s
}
