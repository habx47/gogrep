package parser

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
)

var NoArgErr = errors.New("no arguments provided")
var NoFileArgErr = errors.New("no file arguments provided")
var OpenFileErr = errors.New("cannot open the file")
var ReadFileErr = errors.New("cannot read the file")

func Parser() ([]string, bool, error) {
	var ignoreCase bool
	flag.BoolVar(&ignoreCase, "i", false, "ignore case distinctions")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return nil, ignoreCase, fmt.Errorf("%w", NoArgErr)
	}
	if len(args) == 1 {
		return nil, ignoreCase, fmt.Errorf("%w", NoFileArgErr)
	}

	return args, ignoreCase, nil
}

func GetFileContent(fs fs.FS, file_args []string, search_term string) (string, error) {
	var content bytes.Buffer

	for _, filename := range file_args {
		file, err := fs.Open(filename)
		if err != nil {
			return "", fmt.Errorf("%w: %v", OpenFileErr, err)
		}
		defer file.Close()

		file_content, err := ReadFiles(file)
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}

		content.Write(file_content)
	}
	return content.String(), nil
}

func ReadFiles(r io.Reader) ([]byte, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ReadFileErr, err)
	}
	return content, nil
}
