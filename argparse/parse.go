package argparse

type TokenList []*Token

func (tokenList TokenList) Parse() TokenList {
	previousConversions := 1
	conversions := 1
	tokenList.CheckEndOfOptions()
	for while := true; while; while = previousConversions != 0 && conversions != 0 {
		previousConversions = conversions
		conversions = 0
		for _, token := range tokenList {
			if !token.IsSemantic() {
				token.InferRight()
				if !token.IsSemantic() {
					token.InferLeft()
				}
				if !token.IsSemantic() {
					token.InferPositional()
				}
				if token.IsSemantic() {
					conversions++
				}
			}
		}
	}
	return tokenList
}

func (tokens TokenList) MapToTypes() []TokenType {
	types := make([]TokenType, len(tokens))
	for i := range tokens {
		types[i] = *tokens[i].ttype
	}
	return types
}

func (tokens TokenList) CheckEndOfOptions() {
	for index, token := range tokens {
		switch ttype := (*token.ttype).(type) {
		case ContextFreeTokenType:
		case SemanticTokenType:
			if ttype.Equal(&END_OF_OPTIONS) {
				for rightIndex := index + 1; rightIndex < len(tokens); rightIndex++ {
					(tokens)[rightIndex].setCandidate(&OPERAND)
				}
			}
		}
	}
}

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
		contextFreeTType := *ParseArgument(arg)
		var tokenType TokenType = contextFreeTType
		var semanticCandidates = make([]*SemanticTokenType, len(contextFreeTType.SemanticCandidates))
		copy(semanticCandidates, contextFreeTType.SemanticCandidates)
		token := Token{
			argumentPosition:   i,
			ttype:              &tokenType,
			value:              arg,
			boundTo:            nil,
			semanticCandidates: semanticCandidates,
			tokens:             tokens,
		}
		tokens[i] = &token
	}
	return tokens
}

func ParseArguments(args []string) TokenList {
	tokens := initTokens(args)
	return tokens.Parse()
}
