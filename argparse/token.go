package argparse

import (
	. "cmdse-cli/schema"
	"fmt"
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
	var newCandidates []*SemanticTokenType
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
			if bindings.Contains(semToken.PosModel().Binding) {
				isBound = true
			} else {
				isBound = false
				break
			}
		}
		return isBound
	} else {
		return bindings.Contains(ttype.PosModel().Binding)
	}
}

func (token *Token) IsOptionPart() bool {
	ttype := token.ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isOption := false
		for _, semToken := range token.semanticCandidates {
			if semToken.PosModel().IsOptionPart {
				isOption = true
			} else {
				isOption = false
				break
			}
		}
		return isOption
	} else {
		return ttype.PosModel().IsOptionPart
	}
}
func (token *Token) IsOptionFlag() bool {
	ttype := token.ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isOption := false
		for _, semToken := range token.semanticCandidates {
			if semToken.PosModel().IsOptionFlag {
				isOption = true
			} else {
				isOption = false
				break
			}
		}
		return isOption
	} else {
		return ttype.PosModel().IsOptionFlag
	}
}

func (token *Token) IsSemantic() bool {
	return token.ttype.IsSemantic()
}

func (token *Token) InferLeft() {
	position := token.argumentPosition
	switch token.ttype.(type) {
	case *ContextFreeTokenType:
		if position > 0 {
			leftNeighbour := token.tokens[position-1]
			nbrBoundToLeftOrNone := leftNeighbour.IsBoundToOneOf(Bindings{BindNone, BindLeft})
			if nbrBoundToLeftOrNone {
				if !token.IsOptionPart() {
					// Must be Operand
					token.setCandidate(SemOperand)
				} else {
					// Remove any bound to BindLeft
					token.reduceCandidates(func(tokenType *SemanticTokenType) bool {
						return tokenType.PosModel().Binding != BindLeft
					})
				}
			} else if leftNeighbour.IsBoundTo(BindRight) {
				// Remove any not bound to BindLeft
				token.reduceCandidates(func(tokenType *SemanticTokenType) bool {
					return tokenType.PosModel().Binding == BindLeft
				})
			}
		}
	}
}

func (token *Token) InferRight() {
	position := token.argumentPosition
	switch token.ttype.(type) {
	case *ContextFreeTokenType:
		if token.IsOptionPart() && position < len(token.tokens)+1 {
			rightNeighbour := token.tokens[position+1]
			nbrBoundToRightOrNone := rightNeighbour.IsBoundToOneOf(Bindings{BindNone, BindRight})
			if nbrBoundToRightOrNone {
				// Remove candidates which are bound to right
				token.reduceCandidates(func(tokenType *SemanticTokenType) bool {
					return tokenType.PosModel().Binding != BindRight
				})
			}

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
