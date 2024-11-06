package reader

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc_NewInputReader(t *testing.T) {
	// ARRANGE
	tests := []struct {
		name     string
		args     []string
		expected InputFlags
	}{
		{"with no flag and filepath", []string{"filepath"}, InputFlags{BytesCountFlag: true, LinesCountFlag: true, WordsCountFlag: true, CharsCountFlag: false}},
		{"with count characters flag and filepath", []string{"-m", "filepath"}, InputFlags{BytesCountFlag: false, LinesCountFlag: false, WordsCountFlag: false, CharsCountFlag: true}},
		{"with count words flag and filepath", []string{"-w", "filepath"}, InputFlags{BytesCountFlag: false, LinesCountFlag: false, WordsCountFlag: true, CharsCountFlag: false}},
		{"with count lines and filepath", []string{"-l", "filepath"}, InputFlags{BytesCountFlag: false, LinesCountFlag: true, WordsCountFlag: false, CharsCountFlag: false}},
		{"with count bytes and filepath", []string{"-c", "filepath"}, InputFlags{BytesCountFlag: true, LinesCountFlag: false, WordsCountFlag: false, CharsCountFlag: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir, err := NewInputReader(tt.args)
			defer resetFlags()
			assert.Equal(t, err, nil)
			assert.Equal(t, tt.expected.BytesCountFlag, ir.InputFlags().BytesCountFlag)
			assert.Equal(t, tt.expected.LinesCountFlag, ir.InputFlags().LinesCountFlag)
			assert.Equal(t, tt.expected.WordsCountFlag, ir.InputFlags().WordsCountFlag)
			assert.Equal(t, tt.expected.CharsCountFlag, ir.InputFlags().CharsCountFlag)
		})
	}
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
