package tkn

import (
	"fmt"
	"strings"

	"github.com/cmdse/core/schema"
)

type candidatePredicate func(*schema.SemanticTokenType) bool

type Token struct {
	ArgumentPosition   int
	Ttype              schema.TokenType
	Value              string
	BoundTo            *Token
	Tokens             TokenList
	SemanticCandidates []*schema.SemanticTokenType
}

func NewToken(argumentPosition int, ttype schema.TokenType, value string, tokens TokenList) *Token {
	var semanticCandidates []*schema.SemanticTokenType
	if cfType, ok := ttype.(*schema.ContextFreeTokenType); ok {
		semanticCandidates = make([]*schema.SemanticTokenType, len(cfType.SemanticCandidates))
		copy(semanticCandidates, cfType.SemanticCandidates)
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

func (token *Token) AttemptConvertToSemantic() {
	if len(token.SemanticCandidates) == 1 {
		var semanticType = token.SemanticCandidates[0]
		token.Ttype = semanticType
		if semanticType.PosModel().Binding == schema.BindRight {
			token.BoundTo = token.Tokens[token.ArgumentPosition+1]
		}
		if semanticType.PosModel().Binding == schema.BindLeft {
			token.BoundTo = token.Tokens[token.ArgumentPosition-1]
		}
	}
}

// Remove candidates which don't satisfy to the given candidatePredicate
func (token *Token) reduceCandidates(pred candidatePredicate) {
	newCandidates := make([]*schema.SemanticTokenType, 0, len(token.SemanticCandidates))
	for _, candidate := range token.SemanticCandidates {
		if pred(candidate) {
			newCandidates = append(newCandidates, candidate)
		}
	}
	token.SemanticCandidates = newCandidates
	token.AttemptConvertToSemantic()
}

func (token *Token) setCandidate(tokenType *schema.SemanticTokenType) {
	token.SemanticCandidates = []*schema.SemanticTokenType{tokenType}
	token.AttemptConvertToSemantic()
}

func (token *Token) setCandidates(tokenTypes []*schema.SemanticTokenType) {
	token.SemanticCandidates = tokenTypes
	token.AttemptConvertToSemantic()
}

func (token *Token) IsBoundTo(binding schema.Binding) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(schema.PosModUnset) {
		isBound := true
		for _, semToken := range token.SemanticCandidates {
			if semToken.PosModel().Binding != binding {
				isBound = false
				break
			}
		}
		return isBound
	}
	return ttype.PosModel().Binding == binding
}

func (token *Token) IsBoundToOneOf(bindings schema.Bindings) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(schema.PosModUnset) {
		isBound := false
		for _, semToken := range token.SemanticCandidates {
			isBound = bindings.Contains(semToken.PosModel().Binding)
			if !isBound {
				break
			}
		}
		return isBound
	}
	return bindings.Contains(ttype.PosModel().Binding)
}

// This function returns true if
// - A (token Ttype positional model is Unset) : all its semantic candidates return true for the provided candidatePredicate method
// - B (token type positional model is not Unset) : the candidatePredicate method given token Ttype returns true
func tokenIsWithIndirection(token *Token, predicate func(schema.TokenType) bool) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(schema.PosModUnset) {
		isOption := false
		for _, semToken := range token.SemanticCandidates {
			isOption = predicate(semToken)
			if !isOption {
				break
			}
		}
		return isOption
	}
	return predicate(ttype)
}

// This function returns true if
// - A (token Ttype positional model is Unset) : all its semantic candidates are option parts
// - B (token type positional model is not Unset) : token Ttype is an option part
func (token *Token) IsOptionPart() bool {
	isOptionPart := func(ttype schema.TokenType) bool {
		return ttype.PosModel().IsOptionPart
	}
	return tokenIsWithIndirection(token, isOptionPart)
}

// This function returns true if
// - A (token Ttype positional model is Unset) : all its semantic candidates are option flags
// - B (token type positional model is not Unset) : token Ttype is an option flag
func (token *Token) IsOptionFlag() bool {
	isOptionFlag := func(ttype schema.TokenType) bool {
		return ttype.PosModel().IsOptionFlag
	}
	return tokenIsWithIndirection(token, isOptionFlag)
}

func (token *Token) IsSemantic() bool {
	return token.Ttype.IsSemantic()
}

func (token *Token) IsContextFree() bool {
	return !token.IsSemantic()
}

func inferFromBoundRightLeftNeighbour(token *Token, leftNeighbour *Token) {
	if leftNeighbour.IsBoundTo(schema.BindRight) {
		if semType, ok := leftNeighbour.Ttype.(*schema.SemanticTokenType); ok {
			token.setCandidate(semType.Variant().OptValueTokenType())
		}
	}
}

func inferFromBoundLeftOrNoneLeftNeighbour(token *Token, leftNeighbour *Token) {
	neighbourBoundLeftOrNone := leftNeighbour.IsBoundToOneOf(schema.Bindings{schema.BindNone, schema.BindLeft})
	if neighbourBoundLeftOrNone {
		if !token.IsOptionPart() {
			// Must be Operand
			token.setCandidate(schema.SemOperand)
		} else {
			// Remove any bound to BindLeft
			token.reduceCandidates(func(tokenType *schema.SemanticTokenType) bool {
				return tokenType.PosModel().Binding != schema.BindLeft
			})
		}
	}
}

// This function will return the first non-end-of-options left neighbour if any
func (token *Token) findLeftNeighbour() (neighbour *Token, found bool) {
	position := token.ArgumentPosition
	leftNeighbourPos := position - 1
	hasLeftNeighbour := leftNeighbourPos >= 0
	if hasLeftNeighbour {
		leftNeighbour := token.Tokens[leftNeighbourPos]
		if leftNeighbour.Ttype == schema.SemEndOfOptions {
			return leftNeighbour.findLeftNeighbour()
		} else {
			return leftNeighbour, true
		}
	} else {
		return nil, false
	}
}

func (token *Token) InferLeft() {
	if _, ok := token.Ttype.(*schema.ContextFreeTokenType); ok {
		leftNeighbour, hasLeftNeighbour := token.findLeftNeighbour()
		if hasLeftNeighbour {
			inferFromBoundLeftOrNoneLeftNeighbour(token, leftNeighbour)
			inferFromBoundRightLeftNeighbour(token, leftNeighbour)
		}
	}
}

func inferFromBoundRightOrNoneRightNeighbour(token *Token, rightNeighbour *Token) {
	nbrBoundToRightOrNone := rightNeighbour.IsBoundToOneOf(schema.Bindings{schema.BindNone, schema.BindRight})
	if nbrBoundToRightOrNone {
		// Remove candidates which are bound to right
		token.reduceCandidates(func(tokenType *schema.SemanticTokenType) bool {
			return tokenType.PosModel().Binding != schema.BindRight
		})
	}
}

// This function will return the first non-end-of-options right neighbour if any
func (token *Token) findRightNeighbour() (neighbour *Token, found bool) {
	position := token.ArgumentPosition
	rightNeighbourPos := position + 1
	hasRightNeighbour := rightNeighbourPos < len(token.Tokens)
	if hasRightNeighbour {
		rightNeighbour := token.Tokens[rightNeighbourPos]
		if rightNeighbour.Ttype == schema.SemEndOfOptions {
			return rightNeighbour.findRightNeighbour()
		} else {
			return rightNeighbour, true
		}
	} else {
		return nil, false
	}
}

func (token *Token) InferRight() {
	if _, ok := token.Ttype.(*schema.ContextFreeTokenType); ok {
		rightNeighbour, hasRightNeighbour := token.findRightNeighbour()
		if token.IsOptionPart() && hasRightNeighbour {
			inferFromBoundRightOrNoneRightNeighbour(token, rightNeighbour)
		}
	}
}

func (token *Token) InferPositional() {
	position := token.ArgumentPosition
	if position == len(token.Tokens)-1 {
		if !token.IsOptionPart() {
			token.setCandidate(schema.SemOperand)
		}
	}
}

func (token *Token) ReduceCandidatesWithScheme(scheme schema.OptionScheme) {
	candidates := token.SemanticCandidates
	var newCandidates []*schema.SemanticTokenType
	for _, candidate := range candidates {
		for _, testVariant := range scheme {
			if testVariant == candidate.Variant() || !candidate.PosModel().IsOptionFlag {
				newCandidates = append(newCandidates, candidate)
			}
		}
	}
	token.SemanticCandidates = newCandidates
}

func (token *Token) String() string {
	semCandidateNames := make([]string, len(token.SemanticCandidates))
	for i, candidate := range token.SemanticCandidates {
		semCandidateNames[i] = candidate.String()
	}
	return fmt.Sprintf(`{ 
	pos:%v,
	type:%v,
	Value:'%v',
	BoundTo:%v,
	semCandidates:[%v]
}`, token.ArgumentPosition, token.Ttype, token.Value, token.BoundTo, strings.Join(semCandidateNames, ", "))
}
