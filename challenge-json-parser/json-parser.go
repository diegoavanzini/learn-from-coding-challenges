package jsonparser

type IParser interface {
	Parse(input string) int
}

type Parser struct {
}

// Parse implements IParser.
func (p *Parser) Parse(input string) int {
	parenthesis := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stackToClose := []rune{}
	for _, c := range input {
		if closeP, ok := parenthesis[c]; ok {
			stackToClose = append(stackToClose, closeP)
			continue
		}
		if (c == '}' || c == ']' || c == ')') && (len(stackToClose) == 0 || c != stackToClose[len(stackToClose)-1]) {
			return 1
		} else {
			// parentesi chiusa correttamente
			stackToClose = stackToClose[:len(stackToClose)-1]
		}
	}
	return 0
}

func NewParser() IParser {
	return &Parser{}
}
