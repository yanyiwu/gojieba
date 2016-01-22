package gojieba

/*
#cgo CXXFLAGS: -I./deps -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#include <stdlib.h>
#include "extractor.h"
*/
import "C"

type Extractor struct {
	extractor C.Extractor
}

func NewExtractor(dict_path, hmm_path, user_dict_path, idf_path, stop_word_path string) *Extractor {
	return &Extractor{
		C.NewExtractor(C.CString(dict_path), C.CString(hmm_path), C.CString(user_dict_path), C.CString(idf_path), C.CString(stop_word_path)),
	}
}

func (x *Extractor) Free() {
	C.FreeExtractor(x.extractor)
}

func (x *Extractor) Extract(s string, topk int) []string {
	var words **C.char = C.Extract(x.extractor, C.CString(s), C.int(topk))
	res := cstrings(words)
	C.FreeWords(words)
	return res
}
