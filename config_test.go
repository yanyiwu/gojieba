package gojieba

import (
	"testing"
)

func TestConfig(t *testing.T) {
	file := getCurrentFilePath()
	println(file)
}

func TestGetDictPaths(t *testing.T) {
	println(getDictPaths(DictPath("dictpath"), IdfPath("idfPath"), "user_dict_path"))
}
