package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
