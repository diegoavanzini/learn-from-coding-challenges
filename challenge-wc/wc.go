package main

import "os"

func main() {

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
