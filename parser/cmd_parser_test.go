package parser_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/habx47/gogrep/parser"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i always fail")
}

func TestReadFiles(t *testing.T) {
	t.Run("no error open & read", func(t *testing.T) {
		fs := fstest.MapFS{
			"test.txt": &fstest.MapFile{Data: []byte("hello from test")},
		}
		got, err := parser.ReadFiles(fs, "test.txt")
		want := "hello from test"

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("error on opening", func(t *testing.T) {
		_, err := parser.ReadFiles(StubFailingFS{}, "test.txt")

		if !errors.Is(err, parser.OpenFileErr) {
			t.Errorf("got %v, want %v", err, parser.OpenFileErr)
		}
	})
}
