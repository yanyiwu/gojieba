package gojieba

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleExtract() {
	x := NewExtractor(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
	defer x.Free()

	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	words := x.Extract(s, 5)
	fmt.Println(s)
	fmt.Println("关键词抽取:", strings.Join(words, "/"))

	// Output:
	// 我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。
	// 关键词抽取: CEO/升职/加薪/手扶拖拉机/巅峰
}

func TestExtractor(t *testing.T) {
	x := NewExtractor("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8", "./dict/idf.utf8", "./dict/stop_words.utf8")
	defer x.Free()
	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	expected := "CEO/升职/加薪/手扶拖拉机/巅峰"
	actual := strings.Join(x.Extract(s, 5), "/")
	if expected != actual {
		t.Error(actual)
	}
}

func BenchmarkExtractor(b *testing.B) {
	x := NewExtractor("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8", "./dict/idf.utf8", "./dict/stop_words.utf8")
	defer x.Free()
	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	b.ResetTimer()
	// Stop Timer before x.Free()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		x.Extract(s, 10)
	}
}
