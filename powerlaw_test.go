package ast

import (
	"testing"

	"github.com/GoSym/lexer/tokenizer"
	"github.com/GoSym/parser"
)

func TestPowerlawString(t *testing.T) {
	var powerlaw *PowerLaw
	var base, power Element
	base.data = tokenizer.Token{
		TokenType: tokenizer.RealNumeric,
		Numeric:   0.333333}
	power.data = tokenizer.Token{
		TokenType: tokenizer.Symbol,
		Symbolic:  "π"}
	powerlaw = &PowerLaw{base: &base, power: &power}
	powerlaw = &PowerLaw{base: &base, power: powerlaw}
	new_string := powerlaw.String()
	if new_string != "0.333333^(0.333333^π)" {
		t.Error("Expected \"0.333333^π\", found \"", new_string, "\"")
	}
}

func TestPowerlawTree(t *testing.T) {
	var powerlaw PowerLaw
	var base, power Element
	base.data = tokenizer.Token{
		TokenType: tokenizer.RealNumeric,
		Numeric:   0.333333}
	power.data = tokenizer.Token{
		TokenType: tokenizer.Symbol,
		Symbolic:  "π"}
	operator := parser.MakeSimpleNode(tokenizer.Token{TokenType: tokenizer.Operator, Operator: []rune("•")[0]})
	operator.AddBranch(parser.MakeSimpleNode(base.data))
	operator.AddBranch(parser.MakeSimpleNode(power.data))
	powerlaw = PowerLaw{base: &base, power: &power}
	if (&powerlaw).Tree().Equal(operator) == false {
		t.Error("Expected \"", (&powerlaw).Tree(), "\", found \"", operator, "\"")
	}
}
