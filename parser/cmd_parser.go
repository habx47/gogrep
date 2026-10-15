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

func ReadFiles(fs fs.FS, file_arg string) (string, error) {
	file, err := fs.Open(file_arg)
	if err != nil {
		return "", fmt.Errorf("%w: %v", OpenFileErr, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("%w", ReadFileErr)
	}

	return string(data), nil
}
