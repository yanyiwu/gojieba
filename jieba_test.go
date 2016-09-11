package gojieba

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ExampleJieba() {
	var s string
	var words []string
	use_hmm := true
	//equals with x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	x := NewJieba()
	defer x.Free()

	s = "我来到北京清华大学"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	s = "比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))

	x.AddWord("比特币")
	s = "比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("添加词典后,精确模式:", strings.Join(words, "/"))

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

	s = "区块链"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "长江大桥"
	words = x.CutForSearch(s, !use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	wordinfos := x.Tokenize(s, SearchMode, !use_hmm)
	fmt.Println(s)
	fmt.Println("Tokenize:", wordinfos)

	// Output:
	// 我来到北京清华大学
	// 全模式: 我/来到/北京/清华/清华大学/华大/大学
	// 我来到北京清华大学
	// 精确模式: 我/来到/北京/清华大学
	// 比特币
	// 精确模式: 比特/币
	// 比特币
	// 添加词典后,精确模式: 比特币
	// 他来到了网易杭研大厦
	// 新词识别: 他/来到/了/网易/杭研/大厦
	// 小明硕士毕业于中国科学院计算所，后在日本京都大学深造
	// 搜索引擎模式: 小明/硕士/毕业/于/中国/科学/学院/科学院/中国科学院/计算/计算所/，/后/在/日本/京都/大学/日本京都大学/深造
	// 长春市长春药店
	// 词性标注: 长春市/ns,长春/ns,药店/n
	// 区块链
	// 词性标注: 区块链/nz
	// 长江大桥
	// 搜索引擎模式: 长江/大桥/长江大桥
	// 长江大桥
	// Tokenize: [{长江 0 6} {大桥 6 12} {长江大桥 0 12}]
}

func TestJieba(t *testing.T) {
	//equals with x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	x := NewJieba()
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
	expected = "小明/硕士/毕业/于/中国/科学/学院/科学院/中国科学院/计算/计算所/，/后/在/日本/京都/大学/日本京都大学/深造"
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

	s = "长春市长春药店"
	wordinfos := x.Tokenize(s, SearchMode, false)
	expectedwords := []Word{
		Word{Str: "长春", Start: 0, End: 6},
		Word{Str: "长春市", Start: 0, End: 9},
		Word{Str: "长春", Start: 9, End: 15},
		Word{Str: "药店", Start: 15, End: 21},
	}
	if !reflect.DeepEqual(wordinfos, expectedwords) {
		t.Error()
	}
}

func TestJiebaCutForSearch(t *testing.T) {
	x := NewJieba()
	defer x.Free()
	s := "长江大桥"
	words := x.CutForSearch(s, false)
	expected := []string{
		"长江",
		"大桥",
		"长江大桥",
	}
	if !reflect.DeepEqual(words, expected) {
		t.Error(words, expected)
	}
	wordinfos := x.Tokenize(s, SearchMode, false)
	expectedwords := []Word{
		Word{Str: "长江", Start: 0, End: 6},
		Word{Str: "大桥", Start: 6, End: 12},
		Word{Str: "长江大桥", Start: 0, End: 12},
	}
	if !reflect.DeepEqual(wordinfos, expectedwords) {
		t.Error(wordinfos, expectedwords)
	}
}

func BenchmarkJieba(b *testing.B) {
	//equals with x := NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH)
	x := NewJieba()
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
		x.Tokenize(s, DefaultMode, true)
		x.Tokenize(s, DefaultMode, false)
	}
}

func ExampleExtract() {
	x := NewJieba()
	defer x.Free()

	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	words := x.Extract(s, 5)
	fmt.Println(s)
	fmt.Println("关键词抽取:", strings.Join(words, "/"))
	word_weights := x.ExtractWithWeight(s, 5)
	fmt.Println("关键词抽取:", word_weights)

	// Output:
	// 我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。
	// 关键词抽取: CEO/升职/加薪/手扶拖拉机/巅峰
	// 关键词抽取: [{CEO 11.739204307083542} {升职 10.8561552143} {加薪 10.642581114} {手扶拖拉机 10.0088573539} {巅峰 9.49395840471}]
}

func TestExtractor(t *testing.T) {
	x := NewJieba()
	defer x.Free()
	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	expected := "CEO/升职/加薪/手扶拖拉机/巅峰"
	actual := strings.Join(x.Extract(s, 5), "/")
	if expected != actual {
		t.Error(actual)
	}
}

func BenchmarkExtractor(b *testing.B) {
	// equals with:
	// x := NewExtractor(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
	x := NewJieba()
	defer x.Free()
	s := "我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。"
	b.ResetTimer()
	// Stop Timer before x.Free()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		x.Extract(s, 10)
		x.ExtractWithWeight(s, 10)
	}
}
