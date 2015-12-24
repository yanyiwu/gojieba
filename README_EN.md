# GoJieba [简体中文](README.md)

[![Build Status](https://travis-ci.org/yanyiwu/gojieba.png?branch=master)](https://travis-ci.org/yanyiwu/gojieba) 
[![Author](https://img.shields.io/badge/author-@yanyiwu-blue.svg?style=flat)](http://yanyiwu.com/) 
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)
[![GoDoc](https://godoc.org/github.com/yanyiwu/gojieba?status.svg)](https://godoc.org/github.com/yanyiwu/gojieba)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gojieba/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gojieba?branch=master)

[GoJieba] is a Jieba Chinese Word Segmentation lib written by Go。

## Usage

```
go get github.com/yanyiwu/gojieba
```

See example in [example/demo.go](example/demo.go)

```
cd example
go run demo.go
```

Output Result：

```
全模式: 我/来到/北京/清华/清华大学/华大/大学
精确模式: 我/来到/北京/清华大学
新词识别: 他/来到/了/网易/杭研/大厦
搜索引擎模式: 小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造
```

## Performance

[GoJieba] has a good enough performance,
it maybe is the best of all the Chinese Word Segmentation lib  from the angle of high performance.

Please see more details in [jieba-performance-comparison],
but the article is written by Chinese, Maybe someday it will be transferred to English.

## Contact

```
i@yanyiwu.com
```

[CppJieba]:http://github.com/yanyiwu/cppjieba
[GoJieba]:http://github.com/yanyiwu/gojieba
[jieba-performance-comparison]:http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html
[Jieba]:https://github.com/fxsjy/jieba

