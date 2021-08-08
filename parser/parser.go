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

// Pasrse dem programs
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
