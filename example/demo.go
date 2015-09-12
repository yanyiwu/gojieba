package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
)

func main() {
	x := gojieba.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/user.dict.utf8")
	defer x.Free()
	fmt.Println(x.Cut("南京市长江大桥"))
}
