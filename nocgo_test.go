package gojieba

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildWithoutCgoShowsHelpfulError(t *testing.T) {
	repoRoot, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tmpDir := t.TempDir()
	goMod := `module example.com/nocgo-test

go 1.17

require github.com/yanyiwu/gojieba v0.0.0

replace github.com/yanyiwu/gojieba => ` + repoRoot + `
`
	mainGo := `package main

import "github.com/yanyiwu/gojieba"

var _ *gojieba.Jieba

func main() {}
`

	if err := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(goMod), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(mainGo), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("go", "build", ".")
	cmd.Dir = tmpDir
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0", "GOWORK=off")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected build without cgo to fail")
	}

	out := string(output)
	if !strings.Contains(out, "gojieba requires cgo") || !strings.Contains(out, "CGO_ENABLED=1") {
		t.Fatalf("expected helpful cgo error, got:\n%s", out)
	}
}
