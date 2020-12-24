package gojieba

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	file := getCurrentFilePath()
	println(file)
}

func TestGetDictPaths(t *testing.T) {
	fmt.Println(getDictPaths("", "", "my.user.dict"))
}
