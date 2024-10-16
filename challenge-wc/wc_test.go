package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWc(t *testing.T) {
	wc := NewWc()

	assert.NotNil(t, wc)
}

type Wc struct {
}

func NewWc() Wc {
	return Wc{}
}
