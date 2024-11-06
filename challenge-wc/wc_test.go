package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc_WhenInputFileAndWFlagM_ShouldReturnTheExpectedNumberOfCharacters(t *testing.T) {
	// ARRANGE
	expecteNumberoOfChars := 339292
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	// ACT
	numberOfChars := wc.numberOfChars

	// ASSERT
	assert.Equal(t, expecteNumberoOfChars, numberOfChars)
}

func TestWc_WhenInputFileAndWFlagW_ShouldReturnTheExpectedNumberOfWords(t *testing.T) {
	// ARRANGE
	expecteNumberoOfWords := 58164
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	// ACT
	numberOfWords := wc.numberOfWords

	// ASSERT
	assert.Equal(t, expecteNumberoOfWords, numberOfWords)
}

func TestWc_WhenInputFileAndLFlagL_ShouldReturnTheExpectedNumberOfLines(t *testing.T) {
	// ARRANGE
	expecteNumberoOfRows := 7145
	wc, err := NewWc("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	// ACT
	numberOfRows := wc.numberOfLines

	// ASSERT
	assert.Equal(t, expecteNumberoOfRows, numberOfRows)
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
