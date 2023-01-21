package parser

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joerdav/brev/ast"
	"github.com/joerdav/brev/lexer"
	"github.com/joerdav/brev/tokens"
	"github.com/matryer/is"
)

func TestAssignment(t *testing.T) {
	input := `
	a = 5
	b = 10
	gopher = 838383
	`
	l := lexer.NewLexer(strings.NewReader(input))
	p := New(&l)

	actual := p.ParseProgram()
	expected := &ast.Program{
		Statements: []ast.Statement{
			&ast.AssignmentStatement{
				Token: tokens.Token{
					Type:    tokens.ASSIGN,
					Literal: "=",
					Col:     3,
					Row:     1,
				},
				Name: &ast.Identifier{
					Token: tokens.Token{
						Type:    tokens.IDENT,
						Literal: "a",
						Col:     1,
						Row:     1,
					},
					Value: "a",
				},
				Value: &ast.IntLiteral{
					Token: tokens.Token{
						Type:    tokens.NUMBER,
						Literal: "5",
						Col:     5,
						Row:     1,
					},
					Value: 5,
				},
			},
			&ast.AssignmentStatement{
				Token: tokens.Token{
					Type:    "=",
					Literal: "=",
					Col:     3,
					Row:     2,
				},
				Name: &ast.Identifier{
					Token: tokens.Token{
						Type:    tokens.IDENT,
						Literal: "b",
						Col:     1,
						Row:     2,
					},
					Value: "b",
				},
				Value: &ast.IntLiteral{
					Token: tokens.Token{
						Type:    tokens.NUMBER,
						Literal: "10",
						Col:     5,
						Row:     2,
					},
					Value: 10,
				},
			},
			&ast.AssignmentStatement{
				Token: tokens.Token{
					Type:    "=",
					Literal: "=",
					Col:     8,
					Row:     3,
				},
				Name: &ast.Identifier{
					Token: tokens.Token{
						Type:    tokens.IDENT,
						Literal: "gopher",
						Col:     1,
						Row:     3,
					},
					Value: "gopher",
				},
				Value: &ast.IntLiteral{
					Token: tokens.Token{
						Type:    tokens.NUMBER,
						Literal: "838383",
						Col:     10,
						Row:     3,
					},
					Value: 838383,
				},
			},
		},
	}
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			t.Error(e)
		}
		t.Errorf("got parser errors")
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar"
	l := lexer.NewLexer(strings.NewReader(input))
	p := New(&l)
	actual := p.ParseProgram()
	expected := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token: tokens.Token{Type: tokens.IDENT, Literal: "foobar", Col: 0, Row: 0},
				Expression: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "foobar", Col: 0, Row: 0},
					Value: "foobar",
				},
			},
		},
	}
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			t.Error(e)
		}
		t.Errorf("got parser errors")
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestNumberExpression(t *testing.T) {
	input := "5"
	l := lexer.NewLexer(strings.NewReader(input))
	p := New(&l)
	actual := p.ParseProgram()
	expected := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token: tokens.Token{Type: tokens.NUMBER, Literal: "5", Col: 0, Row: 0},
				Expression: &ast.IntLiteral{
					Token: tokens.Token{Type: tokens.NUMBER, Literal: "5", Col: 0, Row: 0},
					Value: 5,
				},
			},
		},
	}
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			t.Error(e)
		}
		t.Errorf("got parser errors")
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestPrefixExpression(t *testing.T) {
	input := "-5"
	l := lexer.NewLexer(strings.NewReader(input))
	p := New(&l)
	actual := p.ParseProgram()
	expected := &ast.Program{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Token: tokens.Token{Type: tokens.SUB, Literal: "-"},
				Expression: &ast.PrefixExpression{
					Token:    tokens.Token{Type: tokens.SUB, Literal: "-"},
					Operator: "-",
					Right: &ast.IntLiteral{
						Token: tokens.Token{Type: tokens.NUMBER, Literal: "5", Col: 1},
						Value: 5,
					},
				},
			},
		},
	}
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			t.Error(e)
		}
		t.Errorf("got parser errors")
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Fatal(diff)
	}
}

func TestInfixExpressions(t *testing.T) {
	tests := []struct {
		input             string
		leftVal, rightVal int64
		operator          string
	}{
		{"5+5", 5, 5, "+"},
		{"5-5", 5, 5, "-"},
		{"5/5", 5, 5, "/"},
		{"5*5", 5, 5, "*"},
		{"5>5", 5, 5, ">"},
		{"5<5", 5, 5, "<"},
		{"5==5", 5, 5, "=="},
		{"5!=5", 5, 5, "!="},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			is := is.New(t)
			l := lexer.NewLexer(strings.NewReader(tt.input))
			p := New(&l)
			actual := p.ParseProgram()
			if len(p.Errors()) != 0 {
				for _, e := range p.Errors() {
					t.Error(e)
				}
				t.Errorf("got parser errors")
			}
			is.Equal(len(actual.Statements), 1)
			stmt, ok := actual.Statements[0].(*ast.ExpressionStatement)
			is.True(ok)
			exp, ok := stmt.Expression.(*ast.InfixExpression)
			is.True(ok)
			is.Equal(exp.Operator, tt.operator)
			left, ok := exp.Left.(*ast.IntLiteral)
			is.True(ok)
			is.Equal(left.Value, tt.leftVal)
			right, ok := exp.Right.(*ast.IntLiteral)
			is.True(ok)
			is.Equal(right.Value, tt.rightVal)
		})
	}
}
