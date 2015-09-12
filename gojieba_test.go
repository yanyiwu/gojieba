package gojieba

import (
	"reflect"
	"testing"
)

func TestGoJieba(t *testing.T) {
	x := New("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8")
	defer x.Free()

	expected := []string{"南京市", "长江大桥"}
	t.Log(x.Cut("南京市长江大桥"))
	if reflect.DeepEqual(expected, x.Cut("南京市长江大桥")) {
		t.Log("ok")
	} else {
		t.Error("failed.")
	}
}
