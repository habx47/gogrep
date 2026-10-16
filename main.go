package main

import (
	"fmt"
	"log"
	"os"

	"github.com/habx47/gogrep/parser"
)

func main() {
	args, _, err := parser.Parser()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fs := os.DirFS(".")
	file_contents, err := parser.GetFileContent(fs, args[1:], args[0])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("Required content: %v", file_contents)
}
