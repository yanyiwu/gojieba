package gojieba

import (
	"os"
	"path"
)

var (
	DICT_DIR        = path.Join(os.Getenv("GOPATH"), "src/github.com/yanyiwu/gojieba/dict")
	DICT_PATH       = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH        = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH  = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH        = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
)
