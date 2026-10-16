package parser

import (
	"bytes"
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

func GetFileContent(fs fs.FS, file_args []string) (string, error) {
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
