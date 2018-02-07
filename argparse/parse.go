package argparse

import (
	. "cmdse-cli/schema"
)

func ParseArgument(arg string) *ContextFreeTokenType {
	for _, ttype := range ContextFreeTokenTypes {
		regex, isMatchAll := ttype.Regexp()
		if isMatchAll {
			return ttype
		} else {
			if regex.MatchString(arg) {
				return ttype
			}
		}
	}
	return nil
}

func initTokens(args []string) TokenList {
	tokens := make([]*Token, len(args))
	for i, arg := range args {
		contextFreeTType := ParseArgument(arg)
		var tokenType TokenType = contextFreeTType
		var semanticCandidates = make([]*SemanticTokenType, len(contextFreeTType.SemanticCandidates))
		copy(semanticCandidates, contextFreeTType.SemanticCandidates)
		token := Token{
			argumentPosition:   i,
			ttype:              tokenType,
			value:              arg,
			boundTo:            nil,
			semanticCandidates: semanticCandidates,
			tokens:             tokens,
		}
		tokens[i] = &token
		token.possiblyConvertToSemantic()
	}
	return tokens
}

func ParseArguments(args []string) TokenList {
	tokens := initTokens(args)
	return tokens.Parse()
}
