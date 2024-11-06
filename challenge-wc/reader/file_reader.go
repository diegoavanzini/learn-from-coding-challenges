package reader

import "os"

type FileInputReader struct {
	inputFlags InputFlags
	filepath   string
}

func (fir *FileInputReader) Read() ([]byte, error) {
	return os.ReadFile(fir.filepath)
}

func (fir *FileInputReader) InputFlags() InputFlags {
	return fir.inputFlags
}
