package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	byteCountFlag, filepath, err := validateInput(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	wcTool := NewWc(filepath)
	byteCount, err := wcTool.CountBytes()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if *byteCountFlag {
		fmt.Printf("%d %s", byteCount, filepath)
	}
	os.Exit(0)
}

func validateInput(args []string) (byteCount *bool, filepath string, err error) {
	byteCount = flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>")
	flag.CommandLine.Parse(args)

	if !*byteCount {
		return nil, "", errors.New("count what?")
	}
	if len(flag.Args()) == 0 || flag.Args()[0] == "" {
		return nil, "", errors.New("count what? filename is mandatory")
	}
	filepath = flag.Args()[0]
	return byteCount, filepath, nil
}

type wc struct {
	filepath string
}

func (w wc) CountBytes() (int, error) {
	result, err := os.ReadFile(w.filepath)
	return len(result), err
}

func NewWc(inputFile string) wc {
	return wc{
		filepath: inputFile,
	}
}
