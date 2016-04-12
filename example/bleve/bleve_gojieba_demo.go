package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/blevesearch/bleve"
	_ "github.com/yanyiwu/gojieba/bleve"
)

var (
	DICT_DIR       = path.Join(os.Getenv("GOPATH"), "src/github.com/yanyiwu/gojieba/dict")
	DICT_PATH      = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH       = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH = path.Join(DICT_DIR, "user.dict.utf8")
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
			"dictpath":     DICT_PATH,
			"hmmpath":      HMM_PATH,
			"userdictpath": USER_DICT_PATH,
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
