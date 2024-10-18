package main

func main() {

}

type wc struct {
	filepath string
}

func (w wc) CountBytes() int {
	return 0
}

func NewWc(inputFile string) wc {
	return wc{
		filepath: inputFile,
	}
}
