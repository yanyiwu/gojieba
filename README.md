# GoJieba [![Build Status](https://travis-ci.org/yanyiwu/gojieba.png?branch=master)](https://travis-ci.org/yanyiwu/gojieba)

GoJieba 是 Jieba 分词的 Go 语言版本分词库。

## 用法

```
go get github.com/yanyiwu/gojieba
```

示例代码请见 example/demo.go

```
cd example
go run demo.go
```

之所以需要先 cd 到 example 目录下，是因为 demo.go 里面有写死的字典相对路径。

输出结果：

```
全模式: 我/来到/北京/清华/清华大学/华大/大学
精确模式: 我/来到/北京/清华大学
新词识别: 他/来到/了/网易/杭研/大厦
搜索引擎模式: 小明/硕士/毕业/于/中国/中国科学院/科学/科学院/学院/计算所/，/后/在/日本/日本京都大学/京都/京都大学/大学/深造
```

## 性能测试

性能不错，因为 [GoJieba] 本身就是封装了 C++ 版本的 [CppJieba] 而成，
对比测试了一下，耗时大概是 [CppJieba] 的 1.2 倍。
鉴于 [CppJieba] 性能还不错(详见[jieba-performance-comparison])，
所以 [GoJieba] 性能还是可以的，
对于讲究性能的地方还是可以试试的。

## 客服

```
i@yanyiwu.com
```

[CppJieba]:http://github.com/yanyiwu/cppjieba
[GoJieba]:http://github.com/yanyiwu/gojieba
[jieba-performance-comparison]:http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html
