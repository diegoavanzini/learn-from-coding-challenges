package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc_WhenInputFileAndLFlag_ShouldReturnTheExpectedNumberOfLines(t *testing.T) {
	// ARRANGE
	expecteNumberoOfRows := 7145
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer resetFlags()
	// ACT
	numberOfRows := wc.CountRows()

	// ASSERT
	assert.Equal(t, expecteNumberoOfRows, numberOfRows)
}

func TestWc_validateInputWithCountFlagAndFilepath_ShouldReturnTrue(t *testing.T) {
	// ARRANGE
	args := []string{"-c", "filepath"}
	// ACT
	countFlag, filepath, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.Nil(t, err)
	assert.True(t, *countFlag)
	assert.Equal(t, filepath, "filepath")
}
func TestWc_validateInputWithCountFlagButNoFilepath_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{"-c"}

	// ACT
	_, _, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what? filename is mandatory", err.Error())
}
func TestWc_validateInputWithoutArguments_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{}

	// ACT
	_, _, err := validateInput(args)
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
	countBytes := wc.CountBytes()

	// ASSERT
	assert.Equal(t, 342190, countBytes)
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
