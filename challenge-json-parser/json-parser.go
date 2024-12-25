package jsonparser

import (
	"errors"
)

type IParser interface {
	Parse(input string) (*Parsed, error)
}

var brackets = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

type Parser struct {
}

// Parse implements IParser.
func (p *Parser) Parse(input string) (*Parsed, error) {
	var result = NewParsed()
	closeBracketStack := []rune{}
	betweenBrackets := ""
	for _, c := range input {
		if closedBracket, isOpenBracket := brackets[c]; isOpenBracket {
			closeBracketStack = append(closeBracketStack, closedBracket)
			continue
		}
		isCloseBracket := c == '}' || c == ']' || c == ')'
		isTheFirstCloseBracket := len(closeBracketStack) == 0
		if isCloseBracket {
			if isTheFirstCloseBracket || c != closeBracketStack[len(closeBracketStack)-1] {
				return nil, errors.New("bracket closed but not open")
			} else {
				// bracket closed correctly
				closeBracketStack = closeBracketStack[:len(closeBracketStack)-1]
			}
			result.parseSingleKeyValue(betweenBrackets)
		}
		betweenBrackets = betweenBrackets + string(c)
	}
	return result, nil
}

func NewParser() IParser {
	return &Parser{}
}
