package ast

import (
	"github.com/GoSym/lexer/tokenizer"
	"github.com/GoSym/parser"
)

type PowerLaw struct {
	base  AST
	power AST
}

func (p *PowerLaw) String() string {
	var tree *parser.Node = p.Tree()
	var power_s string
	var base_s string
	tree = p.power.Tree()
	if tree.IsOperator() == true {
		power_s = "(" + p.power.String() + ")"
	} else if tree.IsRealNumeric() &&
		tree.Numeric() == 1.0 {
		power_s = ""
	} else {
		power_s = p.power.String()
	}
	tree = p.base.Tree()
	if tree.IsOperator() && tree.Operator() != '^' {
		base_s = "(" + p.base.String() + ")"
	} else {
		base_s = p.base.String()
	}
	if power_s == "" {
		return base_s
	}
	return base_s + "^" + power_s
}

func (p *PowerLaw) Tree() *parser.Node {
	operator := parser.MakeSimpleNode(tokenizer.Token{TokenType: tokenizer.Operator, Operator: []rune("•")[0]})
	base := p.base.Tree()
	power := p.power.Tree()
	operator.AddBranch(base)
	operator.AddBranch(power)
	if power.IsRealNumeric() && power.Numeric() == 1.0 {
		operator = base
	}
	return operator
}
