package argparse

type TokenList []*Token

func (tokens TokenList) Parse() TokenList {
	previousConversions := 1
	conversions := 1
	tokens.CheckEndOfOptions()
	for while := true; while; while = previousConversions != 0 && conversions != 0 {
		previousConversions = conversions
		conversions = 0
		for _, token := range tokens {
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
	return tokens
}

func (tokens TokenList) MapToTypes() []TokenType {
	types := make([]TokenType, len(tokens))
	for i := range tokens {
		types[i] = tokens[i].ttype
	}
	return types
}

func (tokens TokenList) CheckEndOfOptions() {
	for index, token := range tokens {
		switch ttype := token.ttype.(type) {
		case *SemanticTokenType:
			if ttype.Equal(&SemEndOfOptions) {
				for rightIndex := index + 1; rightIndex < len(tokens); rightIndex++ {
					(tokens)[rightIndex].setCandidate(&SemOperand)
				}
			}
		}
	}
}
