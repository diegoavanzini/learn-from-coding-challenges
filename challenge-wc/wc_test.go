package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc_validateInputWithCountWordsFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-w", "filepath"}
	// ACT
	wcinput, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.False(t, *wcinput.BytesCountFlag)
	assert.False(t, *wcinput.LinesCountFlag)
	assert.True(t, *wcinput.WordsCountFlag)
	assert.Equal(t, wcinput.Filepath, "filepath")
}

func TestWc_WhenInputFileAndWFlagW_ShouldReturnTheExpectedNumberOfWords(t *testing.T) {
	// ARRANGE
	expecteNumberoOfWords := 58164
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer resetFlags()
	// ACT
	numberOfWords := wc.numberOfWords

	// ASSERT
	assert.Equal(t, expecteNumberoOfWords, numberOfWords)
}

func TestWc_validateInputWithCountLinesAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-l", "filepath"}
	// ACT
	wcinput, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.False(t, *wcinput.BytesCountFlag)
	assert.True(t, *wcinput.LinesCountFlag)
	assert.Equal(t, wcinput.Filepath, "filepath")
}
func TestWc_validateInputWithCountFlagWAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-w", "filepath"}
	// ACT
	wcinput, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *wcinput.WordsCountFlag)
	assert.Equal(t, wcinput.Filepath, "filepath")
}
func TestWc_WhenInputFileAndLFlagL_ShouldReturnTheExpectedNumberOfLines(t *testing.T) {
	// ARRANGE
	expecteNumberoOfRows := 7145
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer resetFlags()
	// ACT
	numberOfRows := wc.numberOfLines

	// ASSERT
	assert.Equal(t, expecteNumberoOfRows, numberOfRows)
}

func TestWc_validateInputWithCountFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-c", "filepath"}
	// ACT
	wcinput, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *wcinput.BytesCountFlag)
	assert.Equal(t, wcinput.Filepath, "filepath")
}
func TestWc_validateInputWithCountFlagButNoFilepath_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{"-c"}

	// ACT
	_, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what? filename is mandatory", err.Error())
}
func TestWc_validateInputWithoutArguments_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{}

	// ACT
	_, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what?", err.Error())
}

func TestWc_WhenInputFilePathIsWrong_ShouldReturnAnError(t *testing.T) {
	// ARRANGE
	wrongFilePath := "./ghost.txt"

	// ACT
	wc, err := NewWc(wrongFilePath)

	// ASSERT
	assert.Nil(t, wc)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: The system cannot find the file specified.", wrongFilePath), err.Error())
}

func TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes(t *testing.T) {
	// ARRANGE
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	// ACT
	countBytes := wc.numberOfBytes

	// ASSERT
	assert.Equal(t, 342190, countBytes)
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
