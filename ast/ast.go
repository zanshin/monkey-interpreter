package ast

import (
	"github.com/zanshin/interpreter/token"
)

// Abstract Syntax Tree

// Everything is a Node, and has a TokenLiteral
type Node interface {
	TokenLiteral() string
}

// Statements are one kind of node
type Statement interface {
	Node
	statementNode()
}

// Expressions are one kind of node
type Expression interface {
	Node
	expressionNode()
}

// A Program is a slice of Statement
type Program struct {
	Statements []Statement
}

// Satisfy the Node interface requirement
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Identifiers have a value
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Let statements have an identifier on the left, followed by an expression
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Return statements have the return keyword and an expression
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
