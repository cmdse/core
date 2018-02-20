package argparse

import (
	"github.com/cmdse/core/argparse/parser"
	"github.com/cmdse/core/argparse/tkn"
	"github.com/cmdse/core/schema"
)

// ParseArguments turn given arguments into a collection of tokens.
// Not all tokens are guaranteed to be semantic, and human control might be necessary
// to assign the right semantic token type.
func ParseArguments(args []string, pim *schema.ProgramInterfaceModel) tkn.TokenList {
	tokens := parser.InitTokens(args)
	parser := parser.NewParser(tokens, pim, parser.ArgParseBehavior)
	return parser.ParseTokens()
}
