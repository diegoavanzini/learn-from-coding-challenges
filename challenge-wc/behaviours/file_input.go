package behaviours

import (
	"os"
)

type FileInputBehaviour struct {
	inputFlags InputFlags
	Filepath   string
}

func (fir *FileInputBehaviour) WriteResult(stats Counters) string {
	return fir.inputFlags.CountersStringResult(stats) + " " + fir.Filepath
}

func (fir *FileInputBehaviour) ReadInput() ([]byte, error) {
	return os.ReadFile(fir.Filepath)
}

func NewFileInputBehaviour(inputFlags InputFlags, filepath string) (WCToolBehaviour, error) {
	return &FileInputBehaviour{inputFlags, filepath}, nil
}
