package jsonparser

import "strings"

func NewParsed() Parsed {
	return Parsed{}
}

type Parsed map[Key]Value

type Key string
type Value string

func (p Parsed) Get(key Key) Value {
	return p[key]
}

func (p Parsed) Put(key Key, value Value) {
	p[key] = value
}

func (result *Parsed) parseMultipleKeyValues(betweenBrackets string) {
	keyValues := strings.Split(betweenBrackets, ",")
	for _, keyValue := range keyValues {
		result.parseSingleKeyValue(keyValue)
	}
}

func (parsed *Parsed) parseSingleKeyValue(betweenBrackets string) {
	keyValues := strings.Split(betweenBrackets, ":")
	if len(keyValues) > 1 {
		parsed.Put(Key(trimVal(keyValues[0])), Value(strings.TrimSpace(keyValues[1])))
	}
}

func trimVal(x string) string {
	return strings.ReplaceAll(strings.TrimSpace(x), "\"", "")
}
