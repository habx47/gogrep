package parser

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
)

var OpenFileErr = errors.New("cannot open the file")
var ReadFileErr = errors.New("cannot read the file")

func parser() ([]string, bool) {
	var ignoreCase bool
	flag.BoolVar(&ignoreCase, "i", false, "ignore case distinctions")

	flag.Parse()

	return flag.Args(), ignoreCase
}

func GetFileContent(fs fs.FS, file_arg string) (string, error) {
	file, err := fs.Open(file_arg)
	if err != nil {
		return "", fmt.Errorf("%w: %v", OpenFileErr, err)
	}
	defer file.Close()

	content, err := ReadFiles(file)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return string(content), nil
}

func ReadFiles(r io.Reader) ([]byte, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ReadFileErr, err)
	}
	return content, nil
}
