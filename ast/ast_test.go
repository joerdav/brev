package ast

import (
	"testing"

	"github.com/joerdav/brev/tokens"
	"github.com/matryer/is"
)

func TestString(t *testing.T) {
	is := is.New(t)
	program := &Program{
		Statements: []Statement{
			&AssignmentStatement{
				Name:  &Identifier{Value: "foo", Token: tokens.Token{Literal: "foo"}},
				Value: &Identifier{Value: "bar", Token: tokens.Token{Literal: "bar"}},
				Token: tokens.Token{Literal: "foo"},
			},
		},
	}
	is.Equal("foo = bar", program.String())
}
