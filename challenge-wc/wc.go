package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]
	wcinput, err := validateInput(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	wcTool, err := NewWc(wcinput.Filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if *wcinput.BytesCountFlag {
		byteCount := wcTool.numberOfBytes
		fmt.Printf("%d bytes %s\n", byteCount, wcinput.Filepath)
	}
	if *wcinput.LinesCountFlag {
		lineCount := wcTool.numberOfLines
		fmt.Printf("%d lines %s\n", lineCount, wcinput.Filepath)
	}
	if *wcinput.WordsCountFlag {
		wordsCount := wcTool.numberOfWords
		fmt.Printf("%d words %s\n", wordsCount, wcinput.Filepath)
	}
	os.Exit(0)
}

type WordCountInput struct {
	WordsCountFlag *bool
	LinesCountFlag *bool
	BytesCountFlag *bool
	CharsCountFlag *bool
	Filepath       string
}

func validateInput(args []string) (wordCountInput WordCountInput, err error) {
	wordCountInput = WordCountInput{
		BytesCountFlag: flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>"),
		LinesCountFlag: flag.CommandLine.Bool("l", false, "count lines in file usage: ccwc -l <file>"),
		WordsCountFlag: flag.CommandLine.Bool("w", false, "count words in file usage: ccwc -l <file>"),
		CharsCountFlag: flag.CommandLine.Bool("m", false, "count characters in file usage: ccwc -l <file>"),
	}
	flag.CommandLine.Parse(args)

	if !*wordCountInput.BytesCountFlag &&
		!*wordCountInput.LinesCountFlag &&
		!*wordCountInput.WordsCountFlag &&
		!*wordCountInput.CharsCountFlag {
		return wordCountInput, errors.New("count what?")
	}
	if len(flag.Args()) == 0 || flag.Args()[0] == "" {
		return wordCountInput, errors.New("count what? filename is mandatory")
	}
	wordCountInput.Filepath = flag.Args()[0]
	return
}

type wc struct {
	numberOfWords int
	numberOfLines int
	numberOfBytes int
	numberOfChars int
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
		w.numberOfChars += utf8.RuneCount(buf[:c])
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
