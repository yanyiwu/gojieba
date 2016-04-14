package gojieba

/*
#cgo CXXFLAGS: -I./deps -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#include <stdlib.h>
#include "extractor.h"
*/
import "C"
import "unsafe"

type Extractor struct {
	extractor C.Extractor
}

func NewExtractor(dict_path, hmm_path, user_dict_path, idf_path, stop_word_path string) *Extractor {
	dpath := C.CString(dict_path)
	defer C.free(unsafe.Pointer(dpath))
	hpath := C.CString(hmm_path)
	defer C.free(unsafe.Pointer(hpath))
	upath := C.CString(user_dict_path)
	defer C.free(unsafe.Pointer(upath))
	ipath := C.CString(idf_path)
	defer C.free(unsafe.Pointer(ipath))
	spath := C.CString(stop_word_path)
	defer C.free(unsafe.Pointer(spath))
	return &Extractor{
		C.NewExtractor(
			dpath,
			hpath,
			upath,
			ipath,
			spath,
		),
	}
}

func (x *Extractor) Free() {
	C.FreeExtractor(x.extractor)
}

func (x *Extractor) Extract(s string, topk int) []string {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	var words **C.char = C.Extract(x.extractor, cstr, C.int(topk))
	res := cstrings(words)
	defer C.FreeWords(words)
	return res
}
