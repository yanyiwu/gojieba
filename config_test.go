package gojieba

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetCurrentFilePath(t *testing.T) {
	got := getCurrentFilePath()
	want := "config_test.go"

	if filepath.Base(got) != want {
		t.Fatalf("expected current file path to end with %q, got %q", want, got)
	}
}

func TestGetDictPathsDefaults(t *testing.T) {
	got := getDictPaths()
	want := [TOTAL_DICT_PATH_NUMBER]string{
		DICT_PATH,
		HMM_PATH,
		USER_DICT_PATH,
		IDF_PATH,
		STOP_WORDS_PATH,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected default dict paths %v, got %v", want, got)
	}
}

func TestGetDictPathsOverridesNonEmptyValues(t *testing.T) {
	got := getDictPaths("dict.custom", "", "user.custom", "", "stop.custom")
	want := [TOTAL_DICT_PATH_NUMBER]string{
		"dict.custom",
		HMM_PATH,
		"user.custom",
		IDF_PATH,
		"stop.custom",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected overridden dict paths %v, got %v", want, got)
	}
}
