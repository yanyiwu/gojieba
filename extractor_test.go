package gojieba

import (
	"strings"
	"testing"
)

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
