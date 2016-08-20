package gojieba

import (
	"os"
	"path"
	"strings"
)

var (
	DICT_DIR        string
	DICT_PATH       string
	HMM_PATH        string
	USER_DICT_PATH  string
	IDF_PATH        string
	STOP_WORDS_PATH string
)

func init() {
	paths := os.Getenv("GOPATH")
	path_list := strings.Split(paths, ":")
	if len(path_list) == 1 {
		path_list = strings.Split(paths, ";")
	}
	for _, p := range path_list {
		DICT_DIR = path.Join(p, "src/github.com/yanyiwu/gojieba/dict")
		if isDirExists(DICT_DIR) {
			break
		}
	}
	DICT_PATH = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
}

const TOTAL_DICT_PATH_NUMBER = 5

func getDictPaths(args ...string) [TOTAL_DICT_PATH_NUMBER]string {
	dicts := [TOTAL_DICT_PATH_NUMBER]string{
		DICT_PATH,
		HMM_PATH,
		USER_DICT_PATH,
		IDF_PATH,
		STOP_WORDS_PATH,
	}
	for i := 0; i < len(args) && i < len(dicts); i++ {
		dicts[i] = args[i]
	}
	return dicts
}
