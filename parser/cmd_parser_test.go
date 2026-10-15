package parser_test

import (
	"github.com/habx47/gogrep/parser"
	"testing"
	"testing/fstest"
)

func TestReadFiles(t *testing.T) {
	fs := fstest.MapFS{
		"test.txt": &fstest.MapFile{Data: []byte("hello from test")},
	}
	got := parser.ReadFiles(fs, "test.txt")
	want := "hello from test"

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
