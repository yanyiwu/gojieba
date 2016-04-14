package gojieba

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleJieba() {
	var s string
	var words []string
	use_hmm := true
	x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	defer x.Free()

	s = "我来到北京清华大学"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	s = "他来到了网易杭研大厦"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("新词识别:", strings.Join(words, "/"))

	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words = x.CutForSearch(s, use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	s = "长春市长春药店"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	// Output:
	// 我来到北京清华大学
	// 全模式: 我/来到/北京/清华/清华大学/华大/大学
	// 我来到北京清华大学
	// 精确模式: 我/来到/北京/清华大学
	// 他来到了网易杭研大厦
	// 新词识别: 他/来到/了/网易/杭研/大厦
	// 小明硕士毕业于中国科学院计算所，后在日本京都大学深造
	// 搜索引擎模式: 小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造
	// 长春市长春药店
	// 词性标注: 长春市/ns,长春/ns,药店/n
}

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

func TestJieba(t *testing.T) {
	x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
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

func BenchmarkJieba(b *testing.B) {
	x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	s := "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	defer x.Free()
	b.ResetTimer()
	// Stop Timer before x.Free()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		x.Cut(s, false)
		x.Cut(s, true)
		x.CutAll(s)
		x.CutForSearch(s, false)
		x.CutForSearch(s, true)
		x.Tag(s)
	}
}
