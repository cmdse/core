package argparse

import (
	"fmt"
	. "github.com/cmdse/core/schema"
	"strings"
)

type predicate func(*SemanticTokenType) bool

type Token struct {
	argumentPosition   int
	ttype              TokenType
	value              string
	boundTo            *Token
	tokens             TokenList
	semanticCandidates []*SemanticTokenType
}

func newToken(argumentPosition int, ttype TokenType, value string, tokens TokenList) *Token {
	semanticCandidates := []*SemanticTokenType{}
	switch nutype := ttype.(type) {
	case *ContextFreeTokenType:
		semanticCandidates = make([]*SemanticTokenType, len(nutype.SemanticCandidates))
		copy(semanticCandidates, nutype.SemanticCandidates)
	}
	return &Token{
		argumentPosition,
		ttype,
		value,
		nil,
		tokens,
		semanticCandidates,
	}

}

func (token *Token) possiblyConvertToSemantic() {
	if len(token.semanticCandidates) == 1 {
		var semanticType = token.semanticCandidates[0]
		token.ttype = semanticType
		if semanticType.PosModel().Binding == BindRight {
			token.boundTo = token.tokens[token.argumentPosition+1]
		}
		if semanticType.PosModel().Binding == BindLeft {
			token.boundTo = token.tokens[token.argumentPosition-1]
		}
	}
}

// Remove candidates which don't satisfy to the given predicate
func (token *Token) reduceCandidates(pred predicate) {
	newCandidates := make([]*SemanticTokenType, 0, len(token.semanticCandidates))
	for _, candidate := range token.semanticCandidates {
		if pred(candidate) {
			newCandidates = append(newCandidates, candidate)
		}
	}
	token.semanticCandidates = newCandidates
	token.possiblyConvertToSemantic()
}

func (token *Token) setCandidate(tokenType *SemanticTokenType) {
	token.semanticCandidates = []*SemanticTokenType{tokenType}
	token.possiblyConvertToSemantic()
}

func (token *Token) setCandidates(tokenTypes []*SemanticTokenType) {
	token.semanticCandidates = tokenTypes
	token.possiblyConvertToSemantic()
}

func (token *Token) IsBoundTo(binding Binding) bool {
	ttype := token.ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isBound := true
		for _, semToken := range token.semanticCandidates {
			if semToken.PosModel().Binding != binding {
				isBound = false
				break
			}
		}
		return isBound
	} else {
		return ttype.PosModel().Binding == binding
	}
}

func (token *Token) IsBoundToOneOf(bindings Bindings) bool {
	ttype := token.ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isBound := false
		for _, semToken := range token.semanticCandidates {
			isBound = bindings.Contains(semToken.PosModel().Binding)
			if !isBound {
				break
			}
		}
		return isBound
	} else {
		return bindings.Contains(ttype.PosModel().Binding)
	}
}

// This function returns true if
// - A (token ttype positional model is Unset) : all its semantic candidates return true for the provided predicate method
// - B (token type positional model is not Unset) : the predicate method given token ttype returns true
func tokenIsWithIndirection(token *Token, predicate func(TokenType) bool) bool {
	ttype := token.ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isOption := false
		for _, semToken := range token.semanticCandidates {
			isOption = predicate(semToken)
			if !isOption {
				break
			}
		}
		return isOption
	} else {
		return predicate(ttype)
	}
}

// This function returns true if
// - A (token ttype positional model is Unset) : all its semantic candidates are option parts
// - B (token type positional model is not Unset) : token ttype is an option part
func (token *Token) IsOptionPart() bool {
	isOptionPart := func(ttype TokenType) bool {
		return ttype.PosModel().IsOptionPart
	}
	return tokenIsWithIndirection(token, isOptionPart)
}

// This function returns true if
// - A (token ttype positional model is Unset) : all its semantic candidates are option flags
// - B (token type positional model is not Unset) : token ttype is an option flag
func (token *Token) IsOptionFlag() bool {
	isOptionFlag := func(ttype TokenType) bool {
		return ttype.PosModel().IsOptionFlag
	}
	return tokenIsWithIndirection(token, isOptionFlag)
}

func (token *Token) IsSemantic() bool {
	return token.ttype.IsSemantic()
}

func (token *Token) IsContextFree() bool {
	return !token.IsSemantic()
}

func inferFromBoundRightLeftNeighbour(token *Token, leftNeighbour *Token) {
	if leftNeighbour.IsBoundTo(BindRight) {
		switch ttype := leftNeighbour.ttype.(type) {
		case *SemanticTokenType:
			token.setCandidate(ttype.Variant().OptValueTokenType())
		}
	}
}

func inferFromBoundLeftOrNoneLeftNeighbour(token *Token, leftNeighbour *Token) {
	neighbourBoundLeftOrNone := leftNeighbour.IsBoundToOneOf(Bindings{BindNone, BindLeft})
	if neighbourBoundLeftOrNone {
		if !token.IsOptionPart() {
			// Must be Operand
			token.setCandidate(SemOperand)
		} else {
			// Remove any bound to BindLeft
			token.reduceCandidates(func(tokenType *SemanticTokenType) bool {
				return tokenType.PosModel().Binding != BindLeft
			})
		}
	}
}

func (token *Token) InferLeft() {
	position := token.argumentPosition
	switch token.ttype.(type) {
	case *ContextFreeTokenType:
		hasLeftNeighbour := position > 0
		if hasLeftNeighbour {
			leftNeighbour := token.tokens[position-1]
			inferFromBoundLeftOrNoneLeftNeighbour(token, leftNeighbour)
			inferFromBoundRightLeftNeighbour(token, leftNeighbour)
		}
	}
}

func inferFromBoundRightOrNoneRightNeighbour(token *Token, rightNeighbour *Token) {
	nbrBoundToRightOrNone := rightNeighbour.IsBoundToOneOf(Bindings{BindNone, BindRight})
	if nbrBoundToRightOrNone {
		// Remove candidates which are bound to right
		token.reduceCandidates(func(tokenType *SemanticTokenType) bool {
			return tokenType.PosModel().Binding != BindRight
		})
	}
}

func (token *Token) InferRight() {
	position := token.argumentPosition
	switch token.ttype.(type) {
	case *ContextFreeTokenType:
		hasRightNeighbour := position < len(token.tokens)+1
		if token.IsOptionPart() && hasRightNeighbour {
			rightNeighbour := token.tokens[position+1]
			inferFromBoundRightOrNoneRightNeighbour(token, rightNeighbour)
		}
	}
}

func (token *Token) InferPositional() {
	position := token.argumentPosition
	if position == len(token.tokens)-1 {
		if !token.IsOptionPart() {
			token.setCandidate(SemOperand)
		}
	}
}

func (token *Token) ReduceCandidatesWithScheme(scheme OptionScheme) {
	candidates := token.semanticCandidates
	var newCandidates []*SemanticTokenType
	for _, candidate := range candidates {
		for _, testVariant := range scheme {
			if testVariant == candidate.Variant() || !candidate.PosModel().IsOptionFlag {
				newCandidates = append(newCandidates, candidate)
			}
		}
	}
	token.semanticCandidates = newCandidates
}

func (token *Token) String() string {
	semCandidateNames := make([]string, len(token.semanticCandidates))
	for i, candidate := range token.semanticCandidates {
		semCandidateNames[i] = candidate.String()
	}
	return fmt.Sprintf(`{ 
	pos:%v,
	type:%v,
	value:'%v',
	boundTo:%v,
	semCandidates:[%v]
}`, token.argumentPosition, token.ttype, token.value, token.boundTo, strings.Join(semCandidateNames, ", "))
}
