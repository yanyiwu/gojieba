# GoJieba [English](README_EN.md)

[![Build Status](https://travis-ci.org/yanyiwu/gojieba.png?branch=master)](https://travis-ci.org/yanyiwu/gojieba) 
[![Author](https://img.shields.io/badge/author-@yanyiwu-blue.svg?style=flat)](http://yanyiwu.com/) 
[![Performance](https://img.shields.io/badge/performance-excellent-brightgreen.svg?style=flat)](http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html) 
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)
[![GoDoc](https://godoc.org/github.com/yanyiwu/gojieba?status.svg)](https://godoc.org/github.com/yanyiwu/gojieba)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gojieba/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gojieba?branch=master)
[![codebeat badge](https://codebeat.co/badges/a336d042-3583-4212-8204-88da4407438e)](https://codebeat.co/projects/github-com-yanyiwu-gojieba)
[![Go Report Card](https://goreportcard.com/badge/yanyiwu/gojieba)](https://goreportcard.com/report/yanyiwu/gojieba)

[![logo](http://7viirv.com1.z0.glb.clouddn.com/GoJieBaLogo-v2.png)](http://yanyiwu.com/work/2015/09/14/c-cpp-go-mix-programming.html)

[GoJieba]是"结巴"中文分词的Golang语言版本。

## 简介

+ 支持多种分词方式，包括:
+ 最大概率模式
+ HMM新词发现模式
+ 搜索引擎模式
+ 全模式
+ 核心算法底层由C++实现，性能高效。
+ 无缝集成到 [bleve] 到进行搜索引擎的中文分词功能。

## 用法

```
go get github.com/yanyiwu/gojieba
```

See details in [example](example/example_test.go)

输出结果：

```
我来到北京清华大学
全模式: 我/来到/北京/清华/清华大学/华大/大学
我来到北京清华大学
精确模式: 我/来到/北京/清华大学
他来到了网易杭研大厦
新词识别: 他/来到/了/网易/杭研/大厦
小明硕士毕业于中国科学院计算所，后在日本京都大学深造
搜索引擎模式: 小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造
长春市长春药店
词性标注: 长春市/ns,长春/ns,药店/n
```

## Bleve 中文分词插件用法

See Example in [example/bleve/example_test.go](example/bleve/example_test.go)

## 性能评测

[Jieba中文分词系列性能评测]

## 测试

Unittest

```
go test ./...
```

Benchmark

```
go test -bench "Jieba" -test.benchtime 10s
go test -bench "Extractor" -test.benchtime 10s
```

## 客服

+ Email: `i@yanyiwu.com`
+ QQ: 64162451
+ WeChat: 
+ ![image](http://7viirv.com1.z0.glb.clouddn.com/5a7d1b5c0d_yanyiwu_personal_qrcodes.jpg)

[CppJieba]:http://github.com/yanyiwu/cppjieba
[GoJieba]:http://github.com/yanyiwu/gojieba
[Jieba]:https://github.com/fxsjy/jieba
[Jieba中文分词系列性能评测]:http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html
[bleve]:https://github.com/blevesearch/bleve

[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/yanyiwu/gojieba/trend.png)](https://bitdeli.com/free "Bitdeli Badge")
