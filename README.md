# Writing an Interpreter in Go
## Introduction
This repository contains my implementation of the Monkey interpreter as described in Thorsten Ball's
"[Writing an Interpreter in Go](https://interpreterbook.com/)" book.

## Tagging
The tags are derived from the chapter and section heading numbers in the book. E.g., `c2.4` means
chapter 2, section 4. The tags are added to the commit that completes that chapter and section.

## Packages
The interpreter contains these packages.

### ast
The Abstract Syntax Tree, or ast, defines the node types that are used constructing the AST.

* Node interface
* Statement interface
* Expression interface
* Program struct
* LetStatement type
* ReturnStatement type
* ExpressionStatement type
* IntegerLiteral type

### lexer

### parser

### token

### repl
A rudimentary REPL.
