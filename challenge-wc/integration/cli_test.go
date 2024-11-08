package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var targetDir string

type CommandWithArgs struct {
	Command string
	Args    []string
}

func TestCliArgs(t *testing.T) {
	tests := []struct {
		name     string
		commands []CommandWithArgs
		expected string
	}{
		{"Step One", []CommandWithArgs{CommandWithArgs{Command: path.Join(targetDir, "ccwc.exe"), Args: []string{"-c", "test.txt"}}}, "  342190 test.txt\n"},
		{"Step Two", []CommandWithArgs{CommandWithArgs{Command: path.Join(targetDir, "ccwc.exe"), Args: []string{"-l", "test.txt"}}}, "  7145 test.txt\n"},
		{"Step Three", []CommandWithArgs{CommandWithArgs{Command: path.Join(targetDir, "ccwc.exe"), Args: []string{"-w", "test.txt"}}}, "  58164 test.txt\n"},
		{"Step Four", []CommandWithArgs{CommandWithArgs{Command: path.Join(targetDir, "ccwc.exe"), Args: []string{"-m", "test.txt"}}}, "  339292 test.txt\n"},
		{"Step Five", []CommandWithArgs{CommandWithArgs{Command: path.Join(targetDir, "ccwc.exe"), Args: []string{"test.txt"}}}, "  7145  58164  342190 test.txt\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, err := runBinary(tt.commands)

			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("could not get current dir: %v", err)
	}
	targetDir = filepath.Join(dir, "target\\")
	os.Exit(m.Run())
}

func runBinary(commands []CommandWithArgs) ([]byte, error) {
	cmd := exec.Command(commands[0].Command, commands[0].Args...)
	return cmd.CombinedOutput()
}
