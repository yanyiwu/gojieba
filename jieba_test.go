package gojieba

import (
	"strings"
	"testing"
)

func TestJieba(t *testing.T) {
	x := NewJieba("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8")
	defer x.Free()
	var s string
	var expected string
	var actual string
	var use_hmm = true

	s = "我来到北京清华大学"
	expected = "我/来到/北京/清华/清华大学/华大/大学"
	actual = strings.Join(x.CutAll(s), "/")
	if expected != actual {
		t.Error(actual)
	}

	s = "我来到北京清华大学"
	expected = "我/来到/北京/清华大学"
	actual = strings.Join(x.Cut(s, use_hmm), "/")
	if expected != actual {
		t.Error(actual)
	}

	s = "他来到了网易杭研大厦"
	expected = "他/来到/了/网易/杭研/大厦"
	actual = strings.Join(x.Cut(s, use_hmm), "/")
	if expected != actual {
		t.Error(actual)
	}

	s = "他来到了网易杭研大厦"
	expected = "他/来到/了/网易/杭/研/大厦"
	actual = strings.Join(x.Cut(s, !use_hmm), "/")
	if expected != actual {
		t.Error(actual)
	}

	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	expected = "小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造"
	actual = strings.Join(x.CutForSearch(s, use_hmm), "/")
	if expected != actual {
		t.Error(actual)
	}

	s = "长春市长春药店"
	expected = "长春市/ns,长春/ns,药店/n"
	actual = strings.Join(x.Tag(s), ",")
	if expected != actual {
		t.Error(actual)
	}
}
