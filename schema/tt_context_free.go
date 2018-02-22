package schema

import (
	"regexp"
)

type ContextFreeTokenType struct {
	semanticCandidates []*SemanticTokenType
	regexp             *regexp.Regexp
	name               string
}

func (tokenType *ContextFreeTokenType) IsSemantic() bool {
	return tokenType.PosModel().IsSemantic
}

func (*ContextFreeTokenType) PosModel() *PositionalModel {
	return PosModUnset
}

func (tokenType *ContextFreeTokenType) Name() string {
	return tokenType.name
}

func (tokenType *ContextFreeTokenType) Variant() *OptExpressionVariant {
	return nil
}

func (tokenType *ContextFreeTokenType) String() string {
	return tokenType.Name()
}

func (tokenType *ContextFreeTokenType) Equal(comparedTType TokenType) bool {
	return tokenType.Name() == comparedTType.Name()
}

// SemanticCandidates returns a copy of inner semanticCandidates field.
func (tokenType *ContextFreeTokenType) SemanticCandidates() []*SemanticTokenType {
	var semanticCandidates = make([]*SemanticTokenType, len(tokenType.semanticCandidates))
	copy(semanticCandidates, tokenType.semanticCandidates)
	return semanticCandidates
}

func (tokenType *ContextFreeTokenType) Regexp() (*regexp.Regexp, bool) {
	reg := tokenType.regexp
	isMatchAllRegex := reg == nil
	return reg, isMatchAllRegex
}

var (
	CfGnuExplicitAssignment = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemGNUExplicitAssignment,
		},
		regexp: RegexGnuExplicitAssignment,
		name:   "CfGnuExplicitAssignment",
	}
	CfX2lktExplicitAssignment = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemX2lktExplicitAssignment,
		},
		regexp: RegexX2lktExplicitAssignment,
		name:   "CfX2lktExplicitAssignment",
	}
	CfX2lktReverseSwitch = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemX2lktReverseSwitch,
		},
		regexp: RegexX2lktReverseSwitch,
		name:   "CfX2lktReverseSwitch",
	}
	CfEndOfOptions = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemEndOfOptions,
		},
		regexp: RegexEndOfOptions,
		name:   "CfEndOfOptions",
	}
	CfOneDashLetter = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemPOSIXShortAssignmentLeftSide,
			SemPOSIXShortSwitch,
		},
		regexp: RegexOneDashLetter,
		name:   "CfOneDashLetter",
	}
	CfPosixShortStickyValue = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemPOSIXShortStickyValue,
		},
		regexp: RegexPosixShortStickyValue,
		name:   "CfPosixShortStickyValue",
	}
	CfOneDashWordAlphaNum = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemPOSIXStackedShortSwitches,
			SemX2lktSwitch,
			SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: RegexOneDashWordAlphaNum,
		name:   "CfOneDashWordAlphaNum",
	}
	CfOneDashWord = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemX2lktSwitch,
			SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: RegexOneDashWord,
		name:   "CfOneDashWord",
	}
	CfTwoDashWord = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemGNUSwitch,
			SemGNUImplicitAssignmentLeftSide,
		},
		regexp: RegexTwoDashWord,
		name:   "CfTwoDashWord",
	}
	CfOptWord = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemOperand,
			SemPOSIXShortAssignmentValue,
			SemGNUImplicitAssignmentValue,
			SemX2lktImplicitAssignmentValue,
			SemHeadlessOption,
		},
		regexp: RegexOptWord,
		name:   "CfOptWord",
	}
	CfWord = &ContextFreeTokenType{
		semanticCandidates: []*SemanticTokenType{
			SemOperand,
			SemPOSIXShortAssignmentValue,
			SemGNUImplicitAssignmentValue,
			SemX2lktImplicitAssignmentValue,
		},
		name: "CfWord",
	}
)

var ContextFreeTokenTypes = []*ContextFreeTokenType{
	CfGnuExplicitAssignment,
	CfX2lktReverseSwitch,
	CfX2lktExplicitAssignment,
	CfPosixShortStickyValue,
	CfOneDashWordAlphaNum,
	CfEndOfOptions,
	CfTwoDashWord,
	CfOneDashLetter,
	CfOneDashWord,
	CfOptWord,
	CfWord,
}
