package parser_test

import (
	"bytes"
	"errors"
	"io/fs"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
	"testing/iotest"

	"github.com/habx47/gogrep/parser"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("this will always fail on open")
}

func TestGetFileContent(t *testing.T) {
	t.Run("no error open & read", func(t *testing.T) {
		var buffer bytes.Buffer
		fs := fstest.MapFS{
			"test_one.txt": &fstest.MapFile{Data: []byte("hello from test_one\nhello again!\n")},
			"test_two.txt": &fstest.MapFile{Data: []byte("hello from test_two\n")},
		}
		err := parser.GetFileContent(&buffer, fs, []string{"from", "test_one.txt", "test_two.txt"})

		got := buffer.String()
		want := "hello from test_one\nhello from test_two\n"

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("error on opening", func(t *testing.T) {
		var buffer bytes.Buffer
		err := parser.GetFileContent(&buffer, StubFailingFS{}, []string{"", "test.txt"})

		if !errors.Is(err, parser.OpenFileErr) {
			t.Errorf("got %v, want %v", err, parser.OpenFileErr)
		}
	})
}

func TestReadFiles(t *testing.T) {
	t.Run("no error on read", func(t *testing.T) {
		file_content := strings.NewReader("test content")

		got, _ := parser.ReadFiles(file_content)
		want := []byte("test content")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("error on read", func(t *testing.T) {
		simulatedErr := errors.New("this will always fail on read")
		file_content_err := iotest.ErrReader(simulatedErr)

		_, err := parser.ReadFiles(file_content_err)

		if err == nil {
			t.Fatalf("expected an err but got nil")
		}

		if !errors.Is(err, parser.ReadFileErr) {
			t.Errorf("got: %v, want: %v", err, parser.ReadFileErr)
		}
	})
}
