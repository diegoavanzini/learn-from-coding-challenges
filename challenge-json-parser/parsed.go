package jsonparser

import "strings"

func NewParsed() *Parsed {
	return &Parsed{
		values: map[string]interface{}{},
	}
}

type Parsed struct {
	values map[string]interface{}
}

func (p *Parsed) Get(key string) interface{} {
	return p.values[key]
}

func (p *Parsed) Put(key string, value interface{}) {
	p.values[key] = value
}

func (parsed *Parsed) parseSingleKeyValue(betweenBrackets string) {
	keyValues := strings.Split(betweenBrackets, ":")
	if len(keyValues) > 1 {
		parsed.Put(TrimVal(keyValues[0]), TrimVal(keyValues[1]))
	}
}

func TrimVal(x string) string {
	return strings.ReplaceAll(strings.ReplaceAll(x, "\"", ""), " ", "")
}
