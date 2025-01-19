package jsonparser

import (
	"fmt"
	"strconv"
	"strings"
)

func NewParsed() Parsed {
	return Parsed{}
}

/*
 * Parsed is the struct resulting by the Parse of the json
 * in input and is basically a map
 */
type Parsed map[Key]Value

/* Key is a single key in json */
type Key string

/* Value is a single value in json is paired with a Key */
type Value interface{}

/*
 * Get return a single value in the json
 * NOTE: Nested objects are not supported
 */
func (p Parsed) Get(key Key) Value {
	return p[key]
}

func (p Parsed) Put(key Key, value Value) {
	p[key] = value
}

func (p Parsed) Equal(p1 Parsed) bool {
	for k, v := range p {
		if p1[k] == nil || p1[k] != v {
			return false
		}
	}
	return true
}

/* the core of the Parse receinve as input string an object (without brackets)
 * and populate the parsed object itself returning maybe an error */
func (result *Parsed) parseMultipleKeyValues(betweenBrackets string) error {
	var isKey, isValue, isString, isObject, isBoolean, isNumber bool
	var key string
	var value Value = ""
	/* read the input characters and loop */
	for indexCurrentChar := 0; indexCurrentChar < len(betweenBrackets); indexCurrentChar++ {
		currentChar := betweenBrackets[indexCurrentChar]
		if currentChar != '"' && currentChar != ' ' && !isKey && !isValue {
			return fmt.Errorf("keys must be quoted")
		}
		/* begin of the key */
		if currentChar == '"' && !isKey && !isValue {
			isKey = true
			continue
		}
		/* end of the key */
		if currentChar == '"' && isKey && !isValue {
			isKey = false
			continue
		}
		/* end of the key and begin of the value */
		if currentChar == ':' && !isKey && !isValue {
			isValue = true
			isKey = false
			key = strings.Trim(key, " ")
			continue
		}
		/* end of the key and begin of the value but colon not found */
		if currentChar != ' ' && currentChar != ':' && !isKey && !isValue {
			return fmt.Errorf("missing a colon but found %b", currentChar)
		}
		/* build the key */
		if isKey {
			key = fmt.Sprintf("%s%s", key, string(currentChar))
			continue
		}
		/* build the value I have to check if is an object, a string, a number, a boolean */
		if isValue {
			/* the parsed Value is NOT EMPTY and is NOT A
			   STRING and a double quote is an ERROR  */
			if currentChar == '"' && !isObject && !isString && value != "" {
				return fmt.Errorf("is not a string")
			}
			/* the parsed value is NOT EMPTY and it starts with "tru" or "fal" so is a BOOLEAN */
			if !isObject && !isString && value != "" &&
				(strings.HasPrefix(value.(string), "tru") || strings.HasPrefix(value.(string), "fal")) {
				isBoolean = true
			}
			/* the parsed value is NOT EMPTY and by exclusion should be a NUMBER */
			if !isObject && !isString && !isBoolean && value != "" {
				isNumber = true
				if strings.HasPrefix(value.(string), "0") {
					return fmt.Errorf("numbers cannot have leading zeroes")
				}
			}
			/* the value is EMPTY and the first character is a double quotes so it's a STRING */
			if currentChar == '"' && value == "" {
				isString = true
				continue
			}
			/* the value is EMPTY and the first character is a curly bracket so it's an OBJECT */
			if currentChar == '{' && !isObject && value == "" {
				isObject = true
			}
			/* I've read (this is the end of) the Value */
			if (isString && currentChar == '"') || /* it's the end of the string */
				isObject || /* it'is an object */
				currentChar == ',' ||
				currentChar == '}' { /* it'is the end of the object */
				isValue = false
				if isObject {
					/* if the read Value is an object I have to read untill the closed bracket
					   and get the index of the closed bracket */
					closedBracketIndex, err := findMyCloseBracketIndex(betweenBrackets[indexCurrentChar:])
					if err != nil {
						return err
					}
					endIndex := indexCurrentChar + closedBracketIndex + 1
					objectToParse := betweenBrackets[indexCurrentChar:endIndex]
					/* I need to recursively parse the object in the Value */
					value, err = NewJsonParser().Parse(strings.Trim(objectToParse, " "))
					if err != nil {
						return err
					}
					indexCurrentChar = endIndex
				}
				if isBoolean {
					/* if the Value read is a boolean value I have to parse the Value */
					if boolValue, ok := value.(bool); ok {
						value = boolValue
					}
				}
				result.Put(Key(key), value)
				key, value = "", ""
				isNumber, isBoolean, isObject, isString, isValue, isKey = false, false, false, false, false, false
				continue
			}
			if currentChar != ' ' && currentChar != '\n' && currentChar != '\t' {
				value = fmt.Sprintf("%s%s", value, string(currentChar))
			}
		}
	}
	if key != "" && value != nil {
		if isBoolean {
			if boolValue, ok := value.(bool); ok {
				value = boolValue
			}
		}
		if isNumber {
			var err error
			value, err = strconv.Atoi(value.(string))
			if err != nil {
				return err
			}
		}
		result.Put(Key(key), value)
	}
	return nil
}

func findMyCloseBracketIndex(input string) (int, error) {
	openBrackets := 0
	for index, c := range input {
		if c == '{' {
			openBrackets++
		}
		if c == '}' {
			openBrackets--
		}
		if openBrackets == 0 {
			return index, nil
		}
	}
	return 0, fmt.Errorf("close bracket not found")
}
