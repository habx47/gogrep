package parser

import (
	"flag"
	"io"
	"io/fs"
)

func Parser() ([]string, bool) {
	var ignoreCase bool
	flag.BoolVar(&ignoreCase, "i", false, "ignore case distinctions")

	flag.Parse()

	return flag.Args(), ignoreCase
}

func ReadFiles(fs fs.FS, file_arg string) string {
	file, _ := fs.Open(file_arg)
	defer file.Close()

	data, _ := io.ReadAll(file)

	return string(data)
}
