package ast

import (
	"bytes"
	"fmt"

	"github.com/joerdav/brev/tokens"
)

type Node interface {
	fmt.Stringer
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

var _ Node = (*Program)(nil)

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, s := range p.Statements {
		buf.WriteString(s.String())
	}
	return buf.String()
}

var _ Statement = (*AssignmentStatement)(nil)

type AssignmentStatement struct {
	Token tokens.Token
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(as.TokenLiteral())
	buf.WriteString(" = ")
	if as.Value != nil {
		buf.WriteString(as.Value.String())
	}
	return buf.String()
}

var _ Expression = (*Identifier)(nil)

type Identifier struct {
	Token tokens.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

var _ Statement = (*ExpressionStatement)(nil)

type ExpressionStatement struct {
	Token      tokens.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

var _ Expression = (*InfixExpression)(nil)

type InfixExpression struct {
	Token       tokens.Token
	Operator    string
	Right, Left Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var buf bytes.Buffer
	buf.WriteString(ie.Left.String())
	buf.WriteString(ie.Operator)
	buf.WriteString(ie.Right.String())
	return buf.String()
}

var _ Expression = (*PrefixExpression)(nil)

type PrefixExpression struct {
	Token    tokens.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var buf bytes.Buffer
	buf.WriteString(pe.Operator)
	buf.WriteString(pe.Right.String())
	return buf.String()
}

var _ Expression = (*IntLiteral)(nil)

type IntLiteral struct {
	Token tokens.Token
	Value int64
}

func (il *IntLiteral) expressionNode()      {}
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string {
	return fmt.Sprint(il.Value)
}
