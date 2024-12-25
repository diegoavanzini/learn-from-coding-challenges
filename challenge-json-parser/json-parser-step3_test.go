package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonHasMultipleValuesShouldReturnAnObject(t *testing.T) {
	input := `{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}`
	p := NewParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, Value("true"), parsed.Get("key1"))
	assert.Equal(t, Value("101"), parsed.Get("key5"))
	assert.Equal(t, Value("null"), parsed.Get("key3"))
}
