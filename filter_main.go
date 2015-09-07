//go:generate ragel -Z lexer.rl
//go:generate -command yacc go tool yacc
//go:generate yacc -o parser.go -p filter parser.y

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
func NewFilter(str string) (*Filter, error) {
	lexer := newLex([]byte(str))
	ok := filterParse(lexer)

	if ok == 0 {
		return &Filter{lexer.result}, nil
	}
	return nil, errors.New(lexer.err)
}
