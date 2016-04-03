package bleve

import (
	"errors"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/yanyiwu/gojieba"
)

type JiebaTokenizer struct {
	handle *gojieba.Jieba
}

func NewJiebaTokenizer(dictpath, hmmpath, userdictpath string) *JiebaTokenizer {
	return &JiebaTokenizer{
		gojieba.NewJieba(dictpath, hmmpath, userdictpath),
	}
}

func (x *JiebaTokenizer) Free() {
	x.handle.Free()
}

func (x *JiebaTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	start := 0
	end := 0
	pos := 1
	words := x.handle.Cut(string(sentence), false)
	for _, word := range words {
		end = start + len(word)
		token := analysis.Token{
			Term:     []byte(word),
			Start:    start,
			End:      end,
			Position: pos,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
		pos++
		start = end
	}
	return result
}

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	dictpath, ok := config["dictpath"].(string)
	if !ok {
		return nil, errors.New("config dictpath not found")
	}
	hmmpath, ok := config["hmmpath"].(string)
	if !ok {
		return nil, errors.New("config hmmpath not found")
	}
	userdictpath, ok := config["userdictpath"].(string)
	if !ok {
		return nil, errors.New("config userdictpath not found")
	}
	return NewJiebaTokenizer(dictpath, hmmpath, userdictpath), nil
}

func init() {
	registry.RegisterTokenizer("gojieba", tokenizerConstructor)
}
