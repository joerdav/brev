package parser

import (
	"fmt"
	"strconv"

	"github.com/joerdav/brev/ast"
	"github.com/joerdav/brev/lexer"
	"github.com/joerdav/brev/tokens"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[tokens.TokenType]int{
	tokens.EQ:       EQUALS,
	tokens.NE:       EQUALS,
	tokens.LT:       LESSGREATER,
	tokens.GT:       LESSGREATER,
	tokens.ADD:      SUM,
	tokens.SUB:      SUM,
	tokens.SLASH:    PRODUCT,
	tokens.ASTERISK: PRODUCT,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type ParserError struct {
	Message string
	Token   tokens.Token
}

func (pe ParserError) String() string {
	return fmt.Sprintf("%s (line: %d col: %d)", pe.Message, pe.Token.Row, pe.Token.Col)
}

type Parser struct {
	l *lexer.Lexer

	curToken  tokens.Token
	peekToken tokens.Token

	errors []ParserError

	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFns  map[tokens.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []ParserError{}}
	p.nextToken()
	p.nextToken()
	p.addPrefixParser(tokens.IDENT, p.parseIdentifier)
	p.addPrefixParser(tokens.NUMBER, p.parseIntLiteral)
	p.addPrefixParser(tokens.BANG, p.parsePrefixExpression)
	p.addPrefixParser(tokens.SUB, p.parsePrefixExpression)
	p.addInfixParser(tokens.ADD, p.parseInfixExpression)
	p.addInfixParser(tokens.SUB, p.parseInfixExpression)
	p.addInfixParser(tokens.SLASH, p.parseInfixExpression)
	p.addInfixParser(tokens.ASTERISK, p.parseInfixExpression)
	p.addInfixParser(tokens.EQ, p.parseInfixExpression)
	p.addInfixParser(tokens.NE, p.parseInfixExpression)
	p.addInfixParser(tokens.LT, p.parseInfixExpression)
	p.addInfixParser(tokens.GT, p.parseInfixExpression)
	return p
}

func (p *Parser) addPrefixParser(tokenType tokens.TokenType, fn prefixParseFn) {
	if p.prefixParseFns == nil {
		p.prefixParseFns = map[tokens.TokenType]prefixParseFn{}
	}
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) addInfixParser(tokenType tokens.TokenType, fn infixParseFn) {
	if p.infixParseFns == nil {
		p.infixParseFns = map[tokens.TokenType]infixParseFn{}
	}
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) Errors() []ParserError {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := new(ast.Program)
	program.Statements = []ast.Statement{}

	for p.curToken.Type != tokens.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntLiteral() ast.Expression {
	lit := &ast.IntLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as int", p.curToken.Literal)
		p.errors = append(p.errors, ParserError{Message: msg, Token: p.curToken})
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseStatement() ast.Statement {
	if p.curTokenIs(tokens.IDENT) && p.peekTokenIs(tokens.ASSIGN) {
		return p.parseAssignment()
	}
	return p.parseExpressionStatement()
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
		return nil
	}
	leftExp := prefix()
	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseAssignment() *ast.AssignmentStatement {
	stmt := new(ast.AssignmentStatement)
	if !p.peekTokenIs(tokens.ASSIGN) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	stmt.Token = p.curToken
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

func (p *Parser) noPrefixParseFnError(t tokens.Token) {
	msg := fmt.Sprintf("prefix %s not recognised", t.Type)
	p.errors = append(p.errors, ParserError{Token: t, Message: msg})
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)
	return exp
}

func (p *Parser) peekError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, ParserError{Message: msg, Token: p.peekToken})
}

func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
