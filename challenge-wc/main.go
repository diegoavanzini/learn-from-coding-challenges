package main

import (
	"fmt"
	"github.com/diegoavanzini/learnfromcodechallenges/challenge-wc/behaviours"
	"os"
)

func main() {
	args := os.Args[1:]
	behaviour, err := behaviours.Create(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	wcTool, err := NewWcTool(behaviour)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	outputString := wcTool.ResultToString()
	fmt.Println(outputString)
	os.Exit(0)
}
