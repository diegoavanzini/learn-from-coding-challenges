package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonHasValueShouldReturnAnObject(t *testing.T) {
	input := `{"key": "value"}`
	want := "value"
	p := NewParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	got := parsed.Get("key")
	assert.Equal(t, want, got)
}

func Test_whenJsonHasValueShouldReturnAnObjectWithValues(t *testing.T) {
	input := `{"test": "test1"}`
	want := "test1"
	p := NewParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	got := parsed.Get("test")
	assert.Equal(t, want, got)
}
