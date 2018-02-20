package parser

import (
	"github.com/cmdse/core/argparse/tkn"
	"github.com/cmdse/core/schema"
)

func parseArgument(arg string) *schema.ContextFreeTokenType {
	var fetchedType *schema.ContextFreeTokenType
	for _, ttype := range schema.ContextFreeTokenTypes {
		regex, isMatchAll := ttype.Regexp()
		if isMatchAll {
			fetchedType = ttype
			break
		} else {
			if regex.MatchString(arg) {
				fetchedType = ttype
				break
			}
		}
	}
	return fetchedType
}

func InitTokens(args []string) tkn.TokenList {
	tokenList := make(tkn.TokenList, len(args))
	for i, arg := range args {
		contextFreeTType := parseArgument(arg)
		var tokenType schema.TokenType = contextFreeTType
		token := tkn.NewToken(i, tokenType, arg, tokenList)
		tokenList[i] = token
		token.AttemptConvertToSemantic()
	}
	return tokenList
}
