package main

import (
	"bytes"
	"regexp"
	"unicode/utf8"

	"github.com/diegoavanzini/learnfromcodechallenges/challenge-wc/behaviours"
)

type WcTool struct {
	toolBehaviour behaviours.WCToolBehaviour
	counters      behaviours.Counters
}

func (wc *WcTool) calculate() error {
	content, err := wc.toolBehaviour.ReadInput()
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`\s+`)
	wc.counters.NumberOfWords += len(re.FindAllString(string(content), -1))
	wc.counters.NumberOfLines += bytes.Count(content, []byte{'\n'})
	wc.counters.NumberOfBytes += len(content)
	wc.counters.NumberOfChars += utf8.RuneCount(content)
	return nil
}

func (wc *WcTool) ResultToString() string {
	return wc.toolBehaviour.WriteResult(wc.counters)
}

func NewWcTool(ir behaviours.WCToolBehaviour) (*WcTool, error) {
	wc := &WcTool{
		toolBehaviour: ir,
	}
	err := wc.calculate()
	if err != nil {
		return nil, err
	}
	return wc, nil
}
