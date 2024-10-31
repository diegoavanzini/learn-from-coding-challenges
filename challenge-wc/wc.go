package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	lineCountFlag, byteCountFlag, filepath, err := validateInput(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	wcTool, err := NewWc(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if *byteCountFlag {
		byteCount := wcTool.numberOfBytes
		fmt.Printf("%d bytes %s\n", byteCount, filepath)
	}
	if *lineCountFlag {
		lineCount := wcTool.numberOfLines
		fmt.Printf("%d lines %s\n", lineCount, filepath)
	}
	os.Exit(0)
}

func validateInput(args []string) (linesCount, byteCount *bool, filepath string, err error) {
	byteCount = flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>")
	linesCount = flag.CommandLine.Bool("l", false, "count lines in file usage: ccwc -l <file>")

	flag.CommandLine.Parse(args)

	if !*byteCount && !*linesCount {
		return nil, nil, "", errors.New("count what?")
	}
	if len(flag.Args()) == 0 || flag.Args()[0] == "" {
		return nil, nil, "", errors.New("count what? filename is mandatory")
	}
	filepath = flag.Args()[0]
	return linesCount, byteCount, filepath, nil
}

type wc struct {
	numberOfWords int
	numberOfLines int
	numberOfBytes int
	filepath      string
}

func (w *wc) readFile(inputFile string) error {
	r, err := os.Open(w.filepath)
	if err != nil {
		return err
	}
	buf := make([]byte, 32*1024)
	for {
		c, err := r.Read(buf)
		re := regexp.MustCompile(`\s+`)
		w.numberOfWords += len(re.FindAllString(string(buf[:c]), -1))
		w.numberOfLines += bytes.Count(buf[:c], []byte{'\n'})
		w.numberOfBytes += c
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
	}
}

func NewWc(inputFile string) (*wc, error) {
	w := &wc{
		filepath: inputFile,
	}
	err := w.readFile(inputFile)
	if err != nil {
		return nil, err
	}
	return w, err
}
