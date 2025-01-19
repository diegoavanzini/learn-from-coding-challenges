package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonIsInvalidShouldReturn1(t *testing.T) {
	input := "}"
	p := NewJsonParser()
	got, err := p.Parse(input)
	assert.Nil(t, got)
	assert.NotNil(t, err)
}

func Test_whenJsonIsValidShouldReturn0(t *testing.T) {
	input := "{}"
	p := NewJsonParser()
	got, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, NewParsed())
}
