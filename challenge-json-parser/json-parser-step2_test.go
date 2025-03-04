package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonHasValueShouldReturnAnObject(t *testing.T) {
	input := `{"key": "value"}`
	want := Value("value")
	p := NewJsonParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	got := parsed.Get("key")
	assert.Equal(t, want, got)
}

func Test_whenJsonHasValueShouldReturnAnObjectWithValues(t *testing.T) {
	input := `{"test": "test1"}`
	want := Value("test1")
	p := NewJsonParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	got := parsed.Get("test")
	assert.Equal(t, want, got)
}
