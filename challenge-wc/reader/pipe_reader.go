package reader

import (
	"io"
	"os"
)

type PipeInputReader struct {
	inputFlags InputFlags
}

func (pr *PipeInputReader) Read() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}

func (pr *PipeInputReader) InputFlags() InputFlags {
	return pr.inputFlags
}
