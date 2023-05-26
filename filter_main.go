//go:generate ragel -Z lexer.rl
//go:generate goyacc -o parser.go -p filter parser.y
// HACK: we cant modify yacc directly to pass the state
// this is a very ugly way to modify the file,
// but it works + novody cares because its pre-build
//go:generate sh -c "printf '/\tError(s string)/ insert\n\tContext() Context\n.\nxit\neof' | ex parser.go"

// Package gofilter provides implementation of Wireshark display filter.
package gofilter

import (
	"errors"
)

type Message map[string]interface{}

// Filter stores parsed filter string.
type Filter struct {
	root node
}

// Apply checks all values. Return true if filter pass message.
func (f *Filter) Apply(m Message) bool {
	return f.root.Apply(m)
}

// Create new filter
func (ctx Context) NewFilter(str string) (*Filter, error) {
	lexer := newLex([]byte(str), ctx)
	ok := filterParse(lexer)

	if ok == 0 {
		return &Filter{lexer.result}, nil
	}
	return nil, errors.New(lexer.err)
}
