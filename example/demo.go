package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	DICT_DIR        = path.Join(os.Getenv("GOPATH"), "src/github.com/yanyiwu/gojieba/dict")
	DICT_PATH       = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH        = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH  = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH        = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
)

func DemoJieba() {
	println(DICT_DIR)
	fmt.Println("DemoJieba")
	var words []string
	use_hmm := true
	x := gojieba.NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	defer x.Free()

	words = x.CutAll("我来到北京清华大学")
	fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut("我来到北京清华大学", use_hmm)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	words = x.Cut("他来到了网易杭研大厦", use_hmm)
	fmt.Println("新词识别:", strings.Join(words, "/"))

	words = x.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", use_hmm)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	words = x.Tag("长春市长春药店")
	fmt.Println("词性标注:", strings.Join(words, ","))
}

func DemoExtract() {
	fmt.Println("DemoExtract")
	x := gojieba.NewExtractor(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
	defer x.Free()

	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	words := x.Extract(s, 5)
	fmt.Println(s)
	fmt.Println("关键词抽取:", strings.Join(words, "/"))
}

func main() {
	DemoJieba()
	DemoExtract()
}
