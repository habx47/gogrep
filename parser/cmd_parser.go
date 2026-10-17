package parser

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"regexp"
)

var NoArgErr = errors.New("no arguments provided")
var NoFileArgErr = errors.New("no file arguments provided")
var OpenFileErr = errors.New("cannot open the file")
var ReadFileErr = errors.New("cannot read the file")
var WriteErr = errors.New("failed to write matched line to required output")
var RegExpErr = errors.New("invaid regex expression pattern")

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

func ReadFiles(r io.Reader) ([]byte, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ReadFileErr, err)
	}
	return content, nil
}

func FindMatches(w io.Writer, b []byte, search_term string) error {
	r := bytes.NewReader(b)
	scanner := bufio.NewScanner(r)
	re, err := regexp.Compile(search_term)
	if err != nil {
		return fmt.Errorf("%w: %v", RegExpErr, err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		if re.MatchString(line) {
			_, err := fmt.Fprintln(w, line)
			if err != nil {
				return fmt.Errorf("%w: %v", WriteErr, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func Buf_ReadFileMatch(w io.Writer, file fs.File, re *regexp.Regexp) error {
	// could implement bufio.NewReader and then Reader.ReadSting("\n")
	// will have equivalent behaviour while not being limited to buf size

	scanner := bufio.NewScanner(file)

	file_info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	filename := []byte(" (" + file_info.Name() + ")\n")

	for scanner.Scan() {
		line := scanner.Bytes()
		if re.Match(line) {
			if _, err := w.Write(line); err != nil {
				return fmt.Errorf("%w: %v", WriteErr, err)
			}
			if _, err := w.Write(filename); err != nil {
				return fmt.Errorf("%w: %v", WriteErr, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%w: %v", ReadFileErr, err)
	}

	return nil
}

func GetFileContent(w io.Writer, fs fs.FS, args []string, ignoreCase bool) error {
	search_term := args[0]
	if ignoreCase {
		search_term = "(?i)" + search_term
	}
	re, err := regexp.Compile(search_term)

	if err != nil {
		return fmt.Errorf("%w: %v", RegExpErr, err)
	}
	file_args := args[1:]

	for _, filename := range file_args {
		file, err := fs.Open(filename)
		if err != nil {
			fmt.Fprintf(w, "%v: %v\n", OpenFileErr, err)
			continue
		}

		/*	file_content, err := ReadFiles(file)
			file.Close()
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			matchErr := FindMatches(w, file_content, search_term)
			if matchErr != nil {
				return fmt.Errorf("%w", matchErr)
			} */

		matchErr := Buf_ReadFileMatch(w, file, re)
		file.Close()

		if matchErr == nil {
			continue
		}

		if errors.Is(matchErr, ReadFileErr) {
			fmt.Fprintf(w, "%v: %v\n", ReadFileErr, matchErr)
			continue
		}

		return fmt.Errorf("%w", matchErr)
	}

	return nil
}
