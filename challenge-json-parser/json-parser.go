/* JsonParser -- A very simple parser written following the coding challenges
 * 				 exercise and written against the tests hosted here
 * 				 http://www.json.org/JSON_checker/test.zip (see tests_input
 * 				 folder
 *
 * -----------------------------------------------------------------------
 *
 * Copyright (C) 2024 Diego Avanzini <diego dot avanzini at gmail dot com>
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *  *  Redistributions of source code must retain the above copyright
 *     notice, this list of conditions and the following disclaimer.
 *
 *  *  Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package jsonparser

import "errors"

type Parser interface {
	Parse(input string) (Parsed, error)
}

type JsonParser struct{}

/* Parse read the input as a string and return a Parsed object
 * and mayybe an error .
 */
func (p *JsonParser) Parse(input string) (Parsed, error) {
	if len(input) == 1 {
		return nil, errors.New("malformed json")
	}
	if input[0] != '{' || input[len(input)-1] != '}' {
		return nil, errors.New("a JSON payload should be an object or array, not a string")
	}
	/* the default parsedResult is an empty object */
	var parsedResult = NewParsed()
	for currentIndex, currentChar := range input {
		if currentChar == '{' {
			for reversedIndex := len(input) - 1; reversedIndex > 0; reversedIndex-- {
				if input[reversedIndex] == '}' {
					betweenBrackets := input[currentIndex+1 : reversedIndex]
					err := parsedResult.parseMultipleKeyValues(betweenBrackets)
					return parsedResult, err
				}
			}
			return nil, errors.New("malformed json missed '}'")
		}
	}
	return parsedResult, nil
}

func NewJsonParser() Parser {
	return &JsonParser{}
}
