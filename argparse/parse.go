package argparse

import (
	. "cmdse-cli/schema"
)

func ParseArgument(arg string) *ContextFreeTokenType {
	var fetchedType *ContextFreeTokenType
	for _, ttype := range ContextFreeTokenTypes {
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

func initTokens(args []string) TokenList {
	tokens := make([]*Token, len(args))
	for i, arg := range args {
		contextFreeTType := ParseArgument(arg)
		var tokenType TokenType = contextFreeTType
		token := NewToken(i, tokenType, arg, tokens)
		tokens[i] = token
		token.possiblyConvertToSemantic()
	}
	return tokens
}

func ParseArguments(args []string, programInterfaceModel *ProgramInterfaceModel) TokenList {
	tokens := initTokens(args)
	return tokens.Parse(programInterfaceModel)
}
