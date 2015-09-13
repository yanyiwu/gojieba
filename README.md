# GoJieba

GoJieba 是 Jieba 分词的 Go 语言版本分词库。

## Usage

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

## 客服

```
i@yanyiwu.com
```
