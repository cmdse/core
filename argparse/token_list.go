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

func isOperandOrOptAssignmentValue(stt *SemanticTokenType) bool {
	binding := stt.PosModel().Binding
	return (binding == BindLeft || binding == BindNone) && !stt.PosModel().IsOptionFlag
}

func inferFromEndOfOptions(position int, tokens TokenList) {
	// The first token right after end-of-option could be an option assignment value,
	// so we don't treat it as an operand
	tokens[position+1].reduceCandidates(isOperandOrOptAssignmentValue)
	// The tokens after end-of-option must be operands
	for rightIndex := position + 2; rightIndex < len(tokens); rightIndex++ {
		(tokens)[rightIndex].setCandidate(SemOperand)
	}
}

func (tokens TokenList) CheckEndOfOptions() {
	for position, token := range tokens {
		if token.ttype.Equal(SemEndOfOptions) {
			inferFromEndOfOptions(position, tokens)
		}
	}
}

func contextFreeAndOptionFlag(token *Token) bool {
	return token.IsOptionFlag() && token.IsContextFree()
}

func (tokens TokenList) MatchOptionDescription(descriptions OptDescriptionModel) {
	if descriptions != nil {
		for _, token := range tokens.When(contextFreeAndOptionFlag) {
			types := descriptions.MatchArgument(token.value)
			if types != nil {
				token.setCandidates(types)
			}
		}
	}
}

func (tokens TokenList) When(predicate func(token *Token) bool) TokenList {
	predicatedTokens := make(TokenList, 0, len(tokens))
	for _, token := range tokens {
		if predicate(token) {
			predicatedTokens = append(predicatedTokens, token)
		}
	}
	return predicatedTokens
}

func (tokens TokenList) whenContextFree() TokenList {
	return tokens.When((*Token).IsContextFree)
}
