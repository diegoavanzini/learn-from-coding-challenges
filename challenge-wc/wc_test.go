package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	wc := NewWc(wrongFilePath)

	// ACT
	countBytes, err := wc.CountBytes()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: The system cannot find the file specified.", wrongFilePath), err.Error())
	assert.Equal(t, 0, countBytes)
}

func TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes(t *testing.T) {
	// ARRANGE
	wc := NewWc("./test.txt")

	// ACT
	countBytes, err := wc.CountBytes()
	if err != nil {
		t.Fatal(err)
	}

	// ASSERT
	assert.Equal(t, 342190, countBytes)
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
