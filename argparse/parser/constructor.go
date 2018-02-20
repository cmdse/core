package parser

import (
	"github.com/cmdse/core/argparse/tkn"
	"github.com/cmdse/core/schema"
)

func NewParser(tokens tkn.TokenList, pim *schema.ProgramInterfaceModel, behaviour *Behavior) *Parser {
	return &Parser{
		behaviour,
		pim,
		tokens,
		1,
		1,
	}
}
