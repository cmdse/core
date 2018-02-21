package argparse

import (
	"github.com/cmdse/core/schema"
)

// ParseArguments turn given arguments into a collection of tokens.
// Not all tokens are guaranteed to be semantic, and human control might be necessary
// to assign the right semantic token type.
func ParseArguments(args []string, pim *schema.ProgramInterfaceModel) schema.TokenList {
	tokens := InitTokens(args)
	parser := NewParser(tokens, pim, ArgParseBehavior)
	return parser.ParseTokens()
}
