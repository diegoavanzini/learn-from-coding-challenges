package reader

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type InputReader interface {
	Read() ([]byte, error)
	InputFlags() InputFlags
}

type InputFlags struct {
	WordsCountFlag bool
	LinesCountFlag bool
	BytesCountFlag bool
	CharsCountFlag bool
}

func NewInputReader(args []string) (InputReader, error) {
	inputFlags, err := readInputFlags(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fromPipe, err := isPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if fromPipe {
		return &PipeInputReader{
			inputFlags: inputFlags,
		}, nil
	}
	return &FileInputReader{
		inputFlags: inputFlags,
		filepath:   flag.Args()[0],
	}, nil
}

func readInputFlags(args []string) (inputFlags InputFlags, err error) {
	bytesCountFlag := flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>")
	linesCountFlag := flag.CommandLine.Bool("l", false, "count lines in file usage: ccwc -l <file>")
	wordsCountFlag := flag.CommandLine.Bool("w", false, "count words in file usage: ccwc -l <file>")
	charsCountFlag := flag.CommandLine.Bool("m", false, "count characters in file usage: ccwc -l <file>")

	err = flag.CommandLine.Parse(args)
	if err != nil {
		return inputFlags, err
	}

	inputFlags = InputFlags{
		WordsCountFlag: *wordsCountFlag,
		LinesCountFlag: *linesCountFlag,
		BytesCountFlag: *bytesCountFlag,
		CharsCountFlag: *charsCountFlag,
	}
	err = validateInput(&inputFlags)
	if err != nil {
		return inputFlags, err
	}
	return
}

func validateInput(inputFlags *InputFlags) error {
	if !inputFlags.BytesCountFlag &&
		!inputFlags.LinesCountFlag &&
		!inputFlags.WordsCountFlag &&
		!inputFlags.CharsCountFlag {
		inputFlags.BytesCountFlag = true
		inputFlags.WordsCountFlag = true
		inputFlags.LinesCountFlag = true
	}
	if len(flag.Args()) == 0 || flag.Args()[0] == "" {
		return errors.New("count what? filename is mandatory")
	}
	return nil
}

func isPipe() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	modeChar := fi.Mode() & os.ModeCharDevice
	return modeChar == 0, nil
}
