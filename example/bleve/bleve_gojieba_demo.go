package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	_ "github.com/yanyiwu/gojieba/bleve"
)

func main() {
	messages := []struct {
		Id   string
		Body string
	}{
		{
			Id:   "1",
			Body: "你好",
		},
		{
			Id:   "2",
			Body: "世界",
		},
		{
			Id:   "3",
			Body: "亲口",
		},
		{
			Id:   "4",
			Body: "交代",
		},
	}

	indexMapping := bleve.NewIndexMapping()
	dir := "bleve.jieba"
	os.RemoveAll(dir)

	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     "../../dict/jieba.dict.utf8",
			"hmmpath":      "../../dict/hmm_model.utf8",
			"userdictpath": "../../dict/user.dict.utf8",
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"

	index, err := bleve.New(dir, indexMapping)
	if err != nil {
		panic(err)
	}
	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			panic(err)
		}
	}

	querys := []string{
		"你好世界",
		"亲口交代",
	}

	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		x, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(x))
	}
}
