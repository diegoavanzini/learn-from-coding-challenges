package behaviours

import (
	"io"
	"os"
)

type PipeInputBehaviour struct {
	inputFlags InputFlags
}

func (pr *PipeInputBehaviour) WriteResult(stats Counters) string {
	return pr.inputFlags.CountersStringResult(stats)
}

func (pr *PipeInputBehaviour) ReadInput() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}

func NewPipeInputBehaviour(inputFlags InputFlags) (WCToolBehaviour, error) {
	return &PipeInputBehaviour{inputFlags: inputFlags}, nil
}
