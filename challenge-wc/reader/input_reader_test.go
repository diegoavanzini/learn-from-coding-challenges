package reader

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc_NewInputReaderWithNoFlagAndFilepath_ShouldReturnCountBytesLineAndWords(t *testing.T) {
	// ARRANGE
	args := []string{"filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *ir.InputFlags().BytesCountFlag)
	assert.True(t, *ir.InputFlags().LinesCountFlag)
	assert.True(t, *ir.InputFlags().WordsCountFlag)
	assert.False(t, *ir.InputFlags().CharsCountFlag)
}

func TestWc_NewInputReaderWithCountCharactersFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-m", "filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.False(t, *ir.InputFlags().BytesCountFlag)
	assert.False(t, *ir.InputFlags().LinesCountFlag)
	assert.False(t, *ir.InputFlags().WordsCountFlag)
	assert.True(t, *ir.InputFlags().CharsCountFlag)
}

func TestWc_NewInputReaderWithCountWordsFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-w", "filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.False(t, *ir.InputFlags().BytesCountFlag)
	assert.False(t, *ir.InputFlags().LinesCountFlag)
	assert.True(t, *ir.InputFlags().WordsCountFlag)
}

func TestWc_NewInputReaderWithCountLinesAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-l", "filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.False(t, *ir.InputFlags().BytesCountFlag)
	assert.True(t, *ir.InputFlags().LinesCountFlag)
}
func TestWc_NewInputReaderWithCountFlagWAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-w", "filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *ir.InputFlags().WordsCountFlag)
}

func TestWc_NewInputReaderWithCountFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-c", "filepath"}
	// ACT
	ir, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *ir.InputFlags().BytesCountFlag)
}
func TestWc_NewInputReaderWithCountFlagButNoFilepath_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{"-c"}

	// ACT
	_, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what? filename is mandatory", err.Error())
}
func TestWc_NewInputReaderWithoutArguments_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{}

	// ACT
	_, err := NewInputReader(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what? filename is mandatory", err.Error())
}
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
