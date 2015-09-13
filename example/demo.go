package main

import (
	"fmt"

	"github.com/yanyiwu/gojieba"
)

func main() {
	var words []string
	x := gojieba.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/user.dict.utf8")
	defer x.Free()

	words = x.CutAll("我来到北京清华大学")
	fmt.Println("全模式:", words)

	words = x.Cut("我来到北京清华大学", false)
	fmt.Println("精确模式:", words)

	words = x.Cut("他来到了网易杭研大厦", true)
	fmt.Println("新词识别:", words)

	words = x.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true)
	fmt.Println("搜索引擎模式:", words)
}
