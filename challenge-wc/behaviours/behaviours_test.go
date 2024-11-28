package behaviours

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
			ir, err := readInputFlags(tt.args)
			defer ResetFlags()
			assert.Equal(t, err, nil)
			assert.Equal(t, tt.expected.BytesCountFlag, ir.BytesCountFlag)
			assert.Equal(t, tt.expected.LinesCountFlag, ir.LinesCountFlag)
			assert.Equal(t, tt.expected.WordsCountFlag, ir.WordsCountFlag)
			assert.Equal(t, tt.expected.CharsCountFlag, ir.CharsCountFlag)
		})
	}
}

func ResetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
