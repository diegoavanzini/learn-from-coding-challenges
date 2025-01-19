package jsonparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whenJsonHasMultipleValuesAndInternalObject(t *testing.T) {
	input := `{
  "key": "value",
  "key-n": 101,
  "key-o": {
  	"key1": "value1", 
  	"key2": "value2"
  },
  "key-s": {
  	"key1": "value1", 
  	"key2": "value2"
  }
}`
	p := NewJsonParser()
	parsed, err := p.Parse(input)
	if err != nil {
		t.Fatal(err)
	}
	want, err := NewJsonParser().Parse(`{
		"key1": "value1", 
		"key2": "value2"
	}`)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, want.Equal(parsed.Get("key-o").(Parsed)))
	keyo := parsed.Get("key-o").(Parsed)
	assert.Equal(t, "value1", keyo.Get("key1"))
	assert.Equal(t, "value2", keyo.Get("key2"))

	keys := parsed.Get("key-s").(Parsed)
	assert.Equal(t, "value1", keys.Get("key1"))
	assert.Equal(t, "value2", keys.Get("key2"))
}
