package jsonparser

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name  string
	input string
	want  bool
}

func TestParse_WhenInputFromMultipleFiles_ShouldReturnTheExpectedResult(t *testing.T) {
	// ARRANGE
	tests := []test{}
	for i := 1; i <= 33; i++ {
		name := fmt.Sprintf("fail%d", i)
		input := fmt.Sprintf("tests_input/fail%d.json", i)
		tests = append(tests, test{name, input, true})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := ReadFile(tt.input)
			// ACT
			_, err := NewJsonParser().Parse(input)
			// ASSERT
			assert.Equal(t, tt.want, err != nil)
		})
	}
}

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
