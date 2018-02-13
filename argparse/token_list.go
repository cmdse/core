package argparse

import (
	. "github.com/cmdse/core/schema"
)

type TokenList []*Token

func (tokens TokenList) Parse(pim *ProgramInterfaceModel) TokenList {
	previousConversions := 1
	conversions := 1
	tokens.CheckEndOfOptions()
	lastTwoLoopsResultInConversion := func() bool { return previousConversions != 0 && conversions != 0 }
	if scheme := pim.Scheme(); scheme != nil {
		tokens.ReduceCandidatesWithScheme(scheme)
	}
	if descriptions := pim.Descriptions(); descriptions != nil {
		tokens.MatchOptionDescription(descriptions)
	}
	for while := true; while; while = lastTwoLoopsResultInConversion() {
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

func (tokens TokenList) ReduceCandidatesWithScheme(scheme OptionScheme) {
	for _, token := range tokens {
		token.ReduceCandidatesWithScheme(scheme)
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
	for _, token := range tokens {
		if !token.IsSemantic() && token.IsOptionFlag() {
			ttype := descriptions.MatchArgument(token.value)
			if ttype != nil {
				token.setCandidate(ttype)
			}
		}
	}
}
