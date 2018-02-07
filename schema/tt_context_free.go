package schema

import (
	"regexp"
)

type ContextFreeTokenType struct {
	SemanticCandidates []*SemanticTokenType
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

func (tokenType *ContextFreeTokenType) String() string {
	return tokenType.name
}

func (tokenType *ContextFreeTokenType) Equal(comparedTType TokenType) bool {
	return tokenType.Name() == comparedTType.Name()
}

func (tokenType *ContextFreeTokenType) Regexp() (*regexp.Regexp, bool) {
	reg := tokenType.regexp
	isMatchAllRegex := true
	isMatchAllRegex = reg == nil
	return reg, isMatchAllRegex
}

var (
	CfGnuExplicitAssignment = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemGnuExplicitAssignment,
		},
		regexp: RegexGnuExplicitAssignment,
		name:   "CfGnuExplicitAssignment",
	}
	CfX2lktExplicitAssignment = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemX2lktExplicitAssignment,
		},
		regexp: RegexX2lktExplicitAssignment,
		name:   "CfX2lktExplicitAssignment",
	}
	CfX2lktReverseSwitch = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemX2lktReverseSwitch,
		},
		regexp: RegexX2lktReverseSwitch,
		name:   "CfX2lktReverseSwitch",
	}
	CfEndOfOptions = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemEndOfOptions,
		},
		regexp: RegexEndOfOptions,
		name:   "CfEndOfOptions",
	}
	CfOneDashLetter = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemPosixShortAssignmentLeftSide,
			SemPosixShortSwitch,
		},
		regexp: RegexOneDashLetter,
		name:   "CfOneDashLetter",
	}
	CfPosixShortStickyValue = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemPosixShortStickyValue,
		},
		regexp: RegexPosixShortStickyValue,
		name:   "CfPosixShortStickyValue",
	}
	CfOneDashWordAlphaNum = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemPosixStackedShortSwitches,
			SemX2lktSwitch,
			SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: RegexOneDashWordAlphaNum,
		name:   "CfOneDashWordAlphaNum",
	}
	CfOneDashWord = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemX2lktSwitch,
			SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: RegexOneDashWord,
		name:   "CfOneDashWord",
	}
	CfTwoDashWord = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemGnuSwitch,
			SemGnuImplicitAssignmentLeftSide,
		},
		regexp: RegexTwoDashWord,
		name:   "CfTwoDashWord",
	}
	CfWord = &ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			SemOperand,
			SemPosixShortAssignmentValue,
			SemGnuImplicitAssignmentValue,
			SemX2lktImplicitAssignmentValue,
			SemHeadlessOption,
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
	CfWord,
}
