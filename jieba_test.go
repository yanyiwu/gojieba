package gojieba

import (
	"reflect"
	"testing"
)

func TestJieba(t *testing.T) {
	x := New("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/user.dict.utf8")
	defer x.Free()

	expected := []string{"南京市", "长江大桥"}
	actual := x.Cut("南京市长江大桥", false)
	if reflect.DeepEqual(expected, actual) {
		t.Log("ok")
	} else {
		t.Error(actual)
	}
}
