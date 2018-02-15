package argparse

import (
	. "github.com/cmdse/core/schema"
)

type TokenList []*Token

func (tokens TokenList) ReduceCandidatesWithScheme(scheme OptionScheme) {
	if scheme != nil {
		for _, token := range tokens {
			token.ReduceCandidatesWithScheme(scheme)
		}
	}
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
			if ttype.Equal(SemEndOfOptions) {
				for rightIndex := index + 1; rightIndex < len(tokens); rightIndex++ {
					(tokens)[rightIndex].setCandidate(SemOperand)
				}
			}
		}
	}
}

func (tokens TokenList) MatchOptionDescription(descriptions OptDescriptionModel) {
	if descriptions != nil {
		for _, token := range tokens {
			if !token.IsSemantic() && token.IsOptionFlag() {
				types := descriptions.MatchArgument(token.value)
				if types != nil {
					token.setCandidates(types)
				}
			}
		}
	}
}
