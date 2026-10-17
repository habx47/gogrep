package main

import (
	"log"
	"os"

	"github.com/habx47/gogrep/parser"
)

func main() {
	args, ignoreCase, parserErr := parser.Parser()
	if parserErr != nil {
		log.Fatalf("Error: %v\n", parserErr)
	}

	fs := os.DirFS(".")
	err := parser.GetFileContent(os.Stdout, fs, args, ignoreCase)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
