package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis"
	_ "github.com/blevesearch/bleve/analysis/analyzers/custom_analyzer"
	_ "github.com/blevesearch/bleve/registry"
	_ "github.com/yanyiwu/gojieba/bleveplugin"
)

func main() {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "你好",
		From: "i@yanyiwu.com",
		Body: "世界",
	}

	indexMapping := bleve.NewIndexMapping()
	dir := "bleve.jieba"
	os.RemoveAll(dir)

	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     "../dict/jieba.dict.utf8",
			"hmmpath":      "../dict/hmm_model.utf8",
			"userdictpath": "../dict/user.dict.utf8",
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "gojieba",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"

	index, err := bleve.New(dir, indexMapping)
	if err != nil {
		panic(err)
	}
	index.Index(message.Id, message)

	query := bleve.NewQueryStringQuery("你好世界")
	req := bleve.NewSearchRequest(query)
	res, _ := index.Search(req)
	x, _ := json.Marshal(res)
	fmt.Println(string(x))
}
