//go:generate ragel -Z lexer.rl
//go:generate -command yacc go tool yacc
//go:generate yacc -o parser.go -p filter parser.y

package gofilter

import (
	"errors"
)

type Message map[string]interface{}

type Filter struct {
	root node
}

func (f *Filter) Apply(p Message) bool {
	return f.root.Apply(p)
}

func NewFilter(str string) (*Filter, error) {
	lexer := newLex([]byte(str))
	ok := filterParse(lexer)

	if ok == 0 {
		return &Filter{lexer.result}, nil
	}
	return nil, errors.New(lexer.err)
}
