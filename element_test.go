package ast

import (
	"testing"

	"github.com/GoSym/lexer/tokenizer"
	"github.com/GoSym/parser"
)

func TestElementString(t *testing.T) {
	var element Element
	element.data = tokenizer.Token{
		TokenType: tokenizer.RealNumeric,
		Numeric:   0.333333}
	if (&element).String() != "0.333333" {
		t.Error("Expected \"Test\", found ", element.String())
	}
}

func TestElementTree(t *testing.T) {
	var element Element
	element.data = tokenizer.Token{
		TokenType: tokenizer.RealNumeric,
		Numeric:   0.333333}
	if (&element).Tree().Equal(parser.MakeSimpleNode(element.data)) == false {
		t.Error("Expected \"", parser.MakeSimpleNode(element.data), "\", found \"", (&element).Tree(), "\"")
	}
}
