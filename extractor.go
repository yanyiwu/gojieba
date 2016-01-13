package gojieba

/*
#cgo CXXFLAGS: -I./deps -DLOGGING_LEVEL=WARNING -O3 -Wall
#include <stdlib.h>
#include "extractor.h"
*/
import "C"

type Extractor struct {
	extractor C.Extractor
}
