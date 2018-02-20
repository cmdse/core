package argparse

import (
	"github.com/cmdse/core/argparse/parser"
	"github.com/cmdse/core/argparse/tkn"
	"github.com/cmdse/core/schema"
)

func ParseArguments(args []string, pim *schema.ProgramInterfaceModel) tkn.TokenList {
	tokens := parser.InitTokens(args)
	parser := parser.NewParser(tokens, pim, parser.ArgParseBehaviour)
	return parser.ParseTokens()
}
