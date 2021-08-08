package parser

import (
	"github.com/zanshin/interpreter/ast"
	"github.com/zanshin/interpreter/lexer"
	"github.com/zanshin/interpreter/token"
)

// Our Parser has a lexer, and knows about the current token,
// and the one following
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// Return a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two token, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Move foward one token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Create a Program by parsing tokens into statements
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// for p.curToken.Type != token.EOF {
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Process tokens based on their type
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// Let statement has an IDENT and an ASSIGN
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expression until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Some helpers
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
