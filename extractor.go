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

func NewExtractor(paths ...string) *Extractor {
	dictpaths := getDictPaths(paths...)
	dpath, hpath, upath, ipath, spath := C.CString(dictpaths[0]), C.CString(dictpaths[1]), C.CString(dictpaths[2]), C.CString(dictpaths[3]), C.CString(dictpaths[4])
	defer C.free(unsafe.Pointer(dpath))
	defer C.free(unsafe.Pointer(hpath))
	defer C.free(unsafe.Pointer(upath))
	defer C.free(unsafe.Pointer(ipath))
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
