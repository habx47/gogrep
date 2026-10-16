package main

import (
	"log"
	"os"

	"github.com/habx47/gogrep/parser"
)

func main() {
	args, _, parserErr := parser.Parser()
	if parserErr != nil {
		log.Fatalf("Error: %v\n", parserErr)
	}

	fs := os.DirFS(".")
	err := parser.GetFileContent(os.Stdout, fs, args)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
