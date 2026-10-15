package main

import (
	"flag"
	"fmt"
)

var ignoreCase *bool = flag.Bool("i", false, "ignore case distinctions")

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Printf("There are the flags and arguments: %v & %v\n", *ignoreCase, args)
}
