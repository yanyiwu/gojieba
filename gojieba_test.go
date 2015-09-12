package gojieba

import "testing"

func TestGoJieba(t *testing.T) {
	x := New("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8")
	if "你好/世界" == x.Cut("你好世界") {
		t.Log("ok")
	} else {
		t.Error("failed.")
	}
}
