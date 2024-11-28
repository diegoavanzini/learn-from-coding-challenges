package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/diegoavanzini/learnfromcodechallenges/challenge-wc/behaviours"

	"github.com/stretchr/testify/assert"
)

func TestWc_WhenInputFileAndWFlagM_ShouldReturnTheExpectedNumberOfCharacters(t *testing.T) {
	// ARRANGE
	expecteNumberOfChars := 339292
	args := []string{"-m", "test.txt"}
	ir, err := behaviours.Create(args)
	if err != nil {
		t.Fatal(err)
	}
	wcTool, err := NewWcTool(ir)
	if err != nil {
		t.Fatal(err)
	}
	defer ResetFlags()
	// ACT
	numberOfChars := wcTool.counters.NumberOfChars
	actualResult := wcTool.ResultToString()

	// ASSERT
	assert.Equal(t, expecteNumberOfChars, numberOfChars)
	assert.Equal(t, "  339292 test.txt", actualResult)
}

func TestWc_WhenInputFileAndWFlagW_ShouldReturnTheExpectedNumberOfWords(t *testing.T) {
	// ARRANGE
	expecteNumberOfWords := 58164
	args := []string{"-w", "test.txt"}
	ir, err := behaviours.Create(args)
	if err != nil {
		t.Fatal(err)
	}
	wcTool, err := NewWcTool(ir)
	if err != nil {
		t.Fatal(err)
	}
	defer ResetFlags()
	// ACT
	numberOfWords := wcTool.counters.NumberOfWords

	// ASSERT
	assert.Equal(t, expecteNumberOfWords, numberOfWords)
}

func TestWc_WhenInputFileAndLFlagL_ShouldReturnTheExpectedNumberOfLines(t *testing.T) {
	// ARRANGE
	expecteNumberoOfRows := 7145
	args := []string{"-l", "test.txt"}
	ir, err := behaviours.Create(args)
	if err != nil {
		t.Fatal(err)
	}
	wcTool, err := NewWcTool(ir)
	if err != nil {
		t.Fatal(err)
	}
	defer ResetFlags()
	// ACT
	numberOfLines := wcTool.counters.NumberOfLines

	// ASSERT
	assert.Equal(t, expecteNumberoOfRows, numberOfLines)
}

func TestWc_WhenInputFilePathIsWrong_ShouldReturnAnError(t *testing.T) {
	// ARRANGE
	wrongFilePath := "./ghost.txt"
	args := []string{wrongFilePath}
	ir, err := behaviours.Create(args)
	if err != nil {
		t.Fatal(err)
	}
	defer ResetFlags()
	// ACT
	wcTool, err := NewWcTool(ir)

	// ASSERT
	assert.Nil(t, wcTool)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: The system cannot find the file specified.", wrongFilePath), err.Error())
}

func TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes(t *testing.T) {
	// ARRANGE
	args := []string{"test.txt"}
	ir, err := behaviours.Create(args)
	if err != nil {
		t.Fatal(err)
	}
	wcTool, err := NewWcTool(ir)
	if err != nil {
		t.Fatal(err)
	}
	defer ResetFlags()

	// ACT
	countBytes := wcTool.counters.NumberOfBytes

	// ASSERT
	assert.Equal(t, 342190, countBytes)
}

func ResetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
