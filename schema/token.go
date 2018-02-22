package schema

import (
	"fmt"
	"strings"

	"github.com/onsi/ginkgo"
)

// CandidatePredicate is a predicate given a SemanticTokenType
type CandidatePredicate func(*SemanticTokenType) bool

// Token is a dynamic value which hold information about the underlying argument's semantics
type Token struct {
	ArgumentPosition   int
	Ttype              TokenType
	Value              string
	BoundTo            *Token
	Tokens             TokenList
	SemanticCandidates []*SemanticTokenType
}

// NewToken creates a new token. It ttype is context-free, its SemanticCandidates are assigned to a copy of
// ttype's candidates.
func NewToken(argumentPosition int, ttype TokenType, value string, tokens TokenList) *Token {
	token := &Token{
		argumentPosition,
		nil,
		value,
		nil,
		tokens,
		nil,
	}
	token.initTtype(ttype)
	return token
}

// initTtype sets token's type and, if ContextFree, copy semantic candidates to the named field
func (token *Token) initTtype(ttype TokenType) {
	token.Ttype = ttype
	if cfType, ok := ttype.(*ContextFreeTokenType); ok {
		// copy semantic candidates from token type
		var semanticCandidates = make([]*SemanticTokenType, len(cfType.SemanticCandidates))
		copy(semanticCandidates, cfType.SemanticCandidates)
		token.SemanticCandidates = semanticCandidates
	}
}

// AttemptConvertToSemantic assign the only semantic type left in SemanticCandidates
// if its length is 1, do nothing otherwise.
// When such assignment happen, it will assign a bound value to its neighbour
// depending on its positional model.
func (token *Token) AttemptConvertToSemantic() {
	if len(token.SemanticCandidates) == 1 {
		var semanticType = token.SemanticCandidates[0]
		token.Ttype = semanticType
		if semanticType.PosModel().Binding == BindRight {
			rightNeighbour, _ := token.findRightNeighbour()
			token.BoundTo = rightNeighbour
		}
		if semanticType.PosModel().Binding == BindLeft {
			leftNeighbour, _ := token.findLeftNeighbour()
			token.BoundTo = leftNeighbour
		}
		token.SemanticCandidates = nil
	}
}

// ReduceCandidates restrict semantic candidates to those which don't satisfy the given CandidatePredicate
func (token *Token) ReduceCandidates(pred CandidatePredicate) {
	newCandidates := make([]*SemanticTokenType, 0, len(token.SemanticCandidates))
	for _, candidate := range token.SemanticCandidates {
		if pred(candidate) {
			newCandidates = append(newCandidates, candidate)
		}
	}
	token.SemanticCandidates = newCandidates
	token.AttemptConvertToSemantic()
}

func (token *Token) setCandidate(tokenType *SemanticTokenType) {
	token.SemanticCandidates = []*SemanticTokenType{tokenType}
	token.AttemptConvertToSemantic()
}

func (token *Token) setCandidates(tokenTypes []*SemanticTokenType) {
	token.SemanticCandidates = tokenTypes
	token.AttemptConvertToSemantic()
}

// IsBoundTo returns true if the current token is bound to the given binding.
func (token *Token) IsBoundTo(binding Binding) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isBound := len(token.SemanticCandidates) > 0
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

// IsBoundToOneOf returns true
// * when its positional model is unset, if all of its
// semantic candidates' bindings are contained in the given bindings slice.
// * when its positional model is not unset, if its positional model is contained
// in the provided bindings slice.
func (token *Token) IsBoundToOneOf(bindings Bindings) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(PosModUnset) {
		isBound := len(token.SemanticCandidates) > 0
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
// - A (token Ttype positional model is Unset) : all its semantic candidates return true for the provided CandidatePredicate method
// - B (token type positional model is not Unset) : the CandidatePredicate method given token Ttype returns true
func tokenIsWithIndirection(token *Token, predicate func(TokenType) bool) bool {
	ttype := token.Ttype
	if ttype.PosModel().Equal(PosModUnset) {
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

// IsOptionPart returns true if
// * A (token Ttype positional model is Unset) : all its semantic candidates are option parts
// * B (token type positional model is not Unset) : token Ttype is an option part
func (token *Token) IsOptionPart() bool {
	isOptionPart := func(ttype TokenType) bool {
		return ttype.PosModel().IsOptionPart
	}
	return tokenIsWithIndirection(token, isOptionPart)
}

// IsOptionFlag returns true if
// * A (token Ttype positional model is Unset) : all its semantic candidates are option flags
// * B (token type positional model is not Unset) : token Ttype is an option flag
func (token *Token) IsOptionFlag() bool {
	isOptionFlag := func(ttype TokenType) bool {
		return ttype.PosModel().IsOptionFlag
	}
	return tokenIsWithIndirection(token, isOptionFlag)
}

// IsSemantic return true if this token's type is of type SemanticTokenType
func (token *Token) IsSemantic() bool {
	return token.Ttype.IsSemantic()
}

// IsContextFree return true if this token's type is of type ContextFreeTokenType
func (token *Token) IsContextFree() bool {
	return !token.IsSemantic()
}

func omitBoundToLeftCandidates(tokenType *SemanticTokenType) bool {
	return tokenType.PosModel().Binding != BindLeft
}

func omitBoundToRightCandidates(tokenType *SemanticTokenType) bool {
	return tokenType.PosModel().Binding != BindRight
}

func keepBoundToRightCandidates(tokenType *SemanticTokenType) bool {
	return tokenType.PosModel().Binding == BindRight
}

func inferFromBoundRightNeighbourAtLeft(token *Token, leftNeighbour *Token) {
	if leftNeighbour.IsBoundTo(BindRight) {
		if semType, ok := leftNeighbour.Ttype.(*SemanticTokenType); ok {
			token.setCandidate(semType.Variant().OptValueTokenType())
		}
	}
}

func inferFromBoundLeftOrNoneNeighbourAtLeft(token *Token, leftNeighbour *Token) {
	neighbourBoundLeftOrNone := leftNeighbour.IsBoundToOneOf(Bindings{BindNone, BindLeft})
	if neighbourBoundLeftOrNone {
		if !token.IsOptionPart() {
			// Must be Operand
			token.setCandidate(SemOperand)
		} else {
			// Remove any bound to BindLeft
			token.ReduceCandidates(omitBoundToLeftCandidates)
		}
	}
}

func inferFromNoNeighborAtLeft(token *Token) {
	token.ReduceCandidates(omitBoundToLeftCandidates)
}

// This function will return the first non-end-of-options left neighbour if any
func (token *Token) findLeftNeighbour() (neighbour *Token, found bool) {
	position := token.ArgumentPosition
	leftNeighbourPos := position - 1
	hasLeftNeighbour := leftNeighbourPos >= 0
	if hasLeftNeighbour {
		leftNeighbour := token.Tokens[leftNeighbourPos]
		if leftNeighbour.Ttype == SemEndOfOptions {
			return leftNeighbour.findLeftNeighbour()
		}
		return leftNeighbour, true
	} else {
		inferFromNoNeighborAtLeft(token)
	}
	return nil, false
}

// InferLeft will update semantic candidates given its left-neighbour properties.
// If only one semantic candidate remains, the token's type will be assigned its value.
func (token *Token) InferLeft() {
	if _, ok := token.Ttype.(*ContextFreeTokenType); ok {
		leftNeighbour, hasLeftNeighbour := token.findLeftNeighbour()
		if hasLeftNeighbour {
			inferFromBoundLeftOrNoneNeighbourAtLeft(token, leftNeighbour)
			inferFromBoundRightNeighbourAtLeft(token, leftNeighbour)
		}
	}
}

func inferFromBoundRightOrNoneNeighbourAtRight(token *Token, rightNeighbour *Token) {
	nbrBoundToRightOrNone := rightNeighbour.IsBoundToOneOf(Bindings{BindNone, BindRight})
	if nbrBoundToRightOrNone {
		// Remove candidates which are bound to right
		token.ReduceCandidates(omitBoundToRightCandidates)
	}
}

func inferFromNoNeighborAtRight(token *Token) {
	token.ReduceCandidates(omitBoundToRightCandidates)
}

// This function will return the first non-end-of-options right neighbour if any
func (token *Token) findRightNeighbour() (neighbour *Token, found bool) {
	position := token.ArgumentPosition
	rightNeighbourPos := position + 1
	hasRightNeighbour := rightNeighbourPos < len(token.Tokens)
	if hasRightNeighbour {
		rightNeighbour := token.Tokens[rightNeighbourPos]
		if rightNeighbour.Ttype.Equal(SemEndOfOptions) {
			return rightNeighbour.findRightNeighbour()
		}
		return rightNeighbour, true
	}
	return nil, false
}

// InferRight will update semantic candidates given its right-neighbour properties.
// If only one semantic candidate remains, the token's type will be assigned its value.
func (token *Token) InferRight() {
	if _, ok := token.Ttype.(*ContextFreeTokenType); ok {
		rightNeighbour, hasRightNeighbour := token.findRightNeighbour()
		if token.IsOptionPart() && hasRightNeighbour {
			inferFromBoundRightOrNoneNeighbourAtRight(token, rightNeighbour)
			inferFromBoundLeftNeighbourAtRight(token, rightNeighbour)
		}
		if token.IsOptionFlag() && !hasRightNeighbour {
			inferFromNoNeighborAtRight(token)
		}
	}
}

func inferFromBoundLeftNeighbourAtRight(token *Token, rightNeighbour *Token) {
	if rightNeighbour.IsBoundTo(BindLeft) {
		// Keep candidates bound to right
		fmt.Fprintf(ginkgo.GinkgoWriter, "%v\n", rightNeighbour)
		fmt.Fprintf(ginkgo.GinkgoWriter, "%v\n", rightNeighbour.SemanticCandidates)
		token.ReduceCandidates(keepBoundToRightCandidates)
	}
}

// InferPositional will turn the last token to a SemOperand
func (token *Token) InferPositional() {
	position := token.ArgumentPosition
	if position == len(token.Tokens)-1 {
		if !token.IsOptionPart() {
			token.setCandidate(SemOperand)
		}
	}
}

// ReduceCandidatesWithScheme will restrict the set of token's semantic candidates
// to those which comply with the provided OptionScheme.
func (token *Token) ReduceCandidatesWithScheme(scheme OptionScheme) {
	candidates := token.SemanticCandidates
	var newCandidates []*SemanticTokenType
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
