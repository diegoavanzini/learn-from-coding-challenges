package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
