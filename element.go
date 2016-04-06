package ast

import (
	"fmt"

	"github.com/GoSym/lexer/tokenizer"
	"github.com/GoSym/parser"
)

type Element struct {
	data tokenizer.Token
}

func (e *Element) String() string {
	var e_string string
	switch e.data.TokenType {
	case tokenizer.Symbol:
		e_string = e.data.Symbolic
	case tokenizer.RealNumeric:
		e_string = fmt.Sprintf("%f", e.data.Numeric)
	case tokenizer.ImagNumeric:
		e_string = fmt.Sprintf("%fi", e.data.Numeric)
	}
	return e_string
}

func (e *Element) Tree() *parser.Node {
	return parser.MakeSimpleNode(e.data)
}
