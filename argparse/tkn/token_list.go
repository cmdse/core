package tkn

import (
	"github.com/cmdse/core/schema"
)

// TokenList is a list of tokens
type TokenList []*Token

// ReduceCandidatesWithScheme will invoke Token#ReduceCandidatesWithScheme
// for each context-free token in the token list.
func (tokens TokenList) ReduceCandidatesWithScheme(scheme schema.OptionScheme) {
	if scheme != nil {
		for _, token := range tokens.WhenContextFree() {
			token.ReduceCandidatesWithScheme(scheme)
		}
	}
}

// MapToTypes returns a mapping of tokens' TokenTypes
func (tokens TokenList) MapToTypes() []schema.TokenType {
	types := make([]schema.TokenType, len(tokens))
	for i := range tokens {
		types[i] = tokens[i].Ttype
	}
	return types
}

func isOperandOrOptAssignmentValue(stt *schema.SemanticTokenType) bool {
	binding := stt.PosModel().Binding
	return (binding == schema.BindLeft || binding == schema.BindNone) && !stt.PosModel().IsOptionFlag
}

func inferFromEndOfOptions(position int, tokens TokenList) {
	// The first token right after end-of-option could be an option assignment Value,
	// so we don't treat it as an operand
	tokens[position+1].ReduceCandidates(isOperandOrOptAssignmentValue)
	// The Tokens after end-of-option must be operands
	for rightIndex := position + 2; rightIndex < len(tokens); rightIndex++ {
		(tokens)[rightIndex].setCandidate(schema.SemOperand)
	}
}

// CheckEndOfOptions will update token types for token met after an end-of-options token.
func (tokens TokenList) CheckEndOfOptions() {
	for position, token := range tokens {
		if token.Ttype.Equal(schema.SemEndOfOptions) {
			inferFromEndOfOptions(position, tokens)
		}
	}
}

func contextFreeAndOptionFlag(token *Token) bool {
	return token.IsOptionFlag() && token.IsContextFree()
}

// MatchOptionDescription will update token types by matching each of them against the provided descriptionModel
func (tokens TokenList) MatchOptionDescription(descriptionModel schema.OptDescriptionModel) {
	if descriptionModel != nil {
		for _, token := range tokens.When(contextFreeAndOptionFlag) {
			types := descriptionModel.MatchArgument(token.Value)
			if types != nil {
				token.setCandidates(types)
			}
		}
	}
}

// When filter a token list with the given predicate and returns the filtered slice.
func (tokens TokenList) When(predicate func(token *Token) bool) TokenList {
	predicatedTokens := make(TokenList, 0, len(tokens))
	for _, token := range tokens {
		if predicate(token) {
			predicatedTokens = append(predicatedTokens, token)
		}
	}
	return predicatedTokens
}

// WhenContextFree filters the token lists retaining only those which are context-free.
func (tokens TokenList) WhenContextFree() TokenList {
	return tokens.When((*Token).IsContextFree)
}
