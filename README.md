# GoJieba [English](README_EN.md)

[![Build Status](https://travis-ci.org/yanyiwu/gojieba.png?branch=master)](https://travis-ci.org/yanyiwu/gojieba) 
[![Author](https://img.shields.io/badge/author-@yanyiwu-blue.svg?style=flat)](http://yanyiwu.com/) 
[![Performance](https://img.shields.io/badge/performance-excellent-brightgreen.svg?style=flat)](http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html) 
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)
[![GoDoc](https://godoc.org/github.com/yanyiwu/gojieba?status.svg)](https://godoc.org/github.com/yanyiwu/gojieba)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gojieba/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gojieba?branch=master)

[![logo](http://7viirv.com1.z0.glb.clouddn.com/gojieba.jpg)](http://yanyiwu.com/work/2015/09/14/c-cpp-go-mix-programming.html)

[GoJieba]是"结巴"中文分词的Golang语言版本。

## 简介

+ 支持多种分词方式，包括最大概率模式，HMM新词发现模式，搜索引擎模式，全模式等。
+ 核心算法底层由C++实现，性能高效。

## 用法

```
go get github.com/yanyiwu/gojieba
```

示例代码请见 [example/demo.go](example/demo.go)

```
cd example
go run demo.go
```

之所以需要先 cd 到 example 目录下，是因为 demo.go 里面有写死的字典相对路径。

输出结果：

```
DemoJieba
全模式: 我/来到/北京/清华/清华大学/华大/大学
精确模式: 我/来到/北京/清华大学
新词识别: 他/来到/了/网易/杭研/大厦
搜索引擎模式: 小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造
词性标注: 长春市/ns,长春/ns,药店/n
DemoExtract
我是拖拉机学院手扶拖拉机专业的。不用多久，我就会升职加薪，当上CEO，走上人生巅峰。
关键词抽取: CEO/升职/加薪/手扶拖拉机/巅峰
```

## 性能测试

[Jieba中文分词系列性能评测]

## 客服

+ Email: `i@yanyiwu.com`
+ QQ: 64162451
+ WeChat: ![image](http://7viirv.com1.z0.glb.clouddn.com/5a7d1b5c0d_yanyiwu_personal_qrcodes.jpg)

[CppJieba]:http://github.com/yanyiwu/cppjieba
[GoJieba]:http://github.com/yanyiwu/gojieba
[Jieba]:https://github.com/fxsjy/jieba
[Jieba中文分词系列性能评测]:http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html
