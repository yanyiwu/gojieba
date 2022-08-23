package gojieba

import (
	"github.com/goylord/gojieba/deps/cppjieba"
	"github.com/goylord/gojieba/deps/limonp"
	"github.com/goylord/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
