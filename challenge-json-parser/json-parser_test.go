package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonIsInvalidShouldReturn1(t *testing.T) {
	input := "}"
	want := 1
	p := NewParser()
	got := p.Parse(input)
	assert.Equal(t, want, got)
}

func Test_whenJsonIsValidShouldReturn0(t *testing.T) {
	input := "{}"
	want := 0
	p := NewParser()
	got := p.Parse(input)
	assert.Equal(t, want, got)
}
