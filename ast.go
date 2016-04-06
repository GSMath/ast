package ast

import "github.com/GoSym/parser"

type AST interface {
	String() string
	Tree() *parser.Node
}

type Function struct {
	name string
	fx   func(argument AST) AST
}
