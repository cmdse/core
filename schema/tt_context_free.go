package schema

import (
	"fmt"
	"regexp"
)

type ContextFreeTokenType struct {
	SemanticCandidates []*SemanticTokenType
	regexp             string
	name               string
}

func (tokenType *ContextFreeTokenType) IsSemantic() bool {
	return tokenType.PosModel().IsSemantic
}

func (*ContextFreeTokenType) PosModel() *PositionalModel {
	return &UNSET
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
	// each regex is tested, so compilable regex can be asserted
	reg := regexp.MustCompile(tokenType.regexp)
	isMatchAllRegex := true
	isMatchAllRegex = reg.String() == ""
	return reg, isMatchAllRegex
}

const regexAlphaNumChar = `[A-Za-z0-9]`
const regexAlphaChar = `[A-Za-z]`
const regexNumChar = `[0-9]`
const regexOptionChar = `[\w_\.-]`
const regexValueWordGroup = `(.*)`

var regexOptionWordGroup = fmt.Sprintf(`(%s%s+)`, regexAlphaNumChar, regexOptionChar)

var (
	CfGnuExplicitAssignment = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemGnuExplicitAssignment,
		},
		regexp: fmt.Sprintf(`^--%s=%s$`, regexOptionWordGroup, regexValueWordGroup),
		name:   "CfGnuExplicitAssignment",
	}
	CfX2lktExplicitAssignment = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemX2lktExplicitAssignment,
		},
		regexp: fmt.Sprintf(`^-%s=%s$`, regexOptionWordGroup, regexValueWordGroup),
		name:   "CfX2lktExplicitAssignment",
	}
	CfX2lktReverseSwitch = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemX2lktReverseSwitch,
		},
		regexp: fmt.Sprintf(`^\+%s$`, regexOptionWordGroup),
		name:   "CfX2lktReverseSwitch",
	}
	CfEndOfOptions = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemEndOfOptions,
		},
		regexp: `^--$`,
		name:   "CfEndOfOptions",
	}
	CfOneDashLetter = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemPosixShortAssignmentLeftSide,
			&SemPosixShortSwitch,
		},
		regexp: fmt.Sprintf(`^-(%s)$`, regexAlphaNumChar),
		name:   "CfOneDashLetter",
	}
	CfPosixShortStickyValue = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemPosixShortStickyValue,
		},
		regexp: fmt.Sprintf(`^-(%s)(%s+)$`, regexAlphaChar, regexNumChar),
		name:   "CfPosixShortStickyValue",
	}
	CfOneDashWordAlphaNum = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemPosixStackedShortSwitches,
			&SemX2lktSwitch,
			&SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: fmt.Sprintf(`^-(%s{2,})$`, regexAlphaNumChar),
		name:   "CfOneDashWordAlphaNum",
	}
	CfOneDashWord = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemX2lktSwitch,
			&SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: fmt.Sprintf(`^-%s$`, regexOptionWordGroup),
		name:   "CfOneDashWord",
	}
	CfTwoDashWord = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemGnuSwitch,
			&SemGnuImplicitAssignmentLeftSide,
		},
		regexp: fmt.Sprintf(`^--%s$`, regexOptionWordGroup),
		name:   "CfTwoDashWord",
	}
	CfWord = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemOperand,
			&SemPosixShortAssignmentValue,
			&SemGnuImplicitAssignmentValue,
			&SemX2lktImplicitAssignmentValue,
			&SemHeadlessOption,
		},
		name: "CfWord",
	}
)

var ContextFreeTokenTypes = []*ContextFreeTokenType{
	&CfGnuExplicitAssignment,
	&CfX2lktReverseSwitch,
	&CfX2lktExplicitAssignment,
	&CfPosixShortStickyValue,
	&CfOneDashWordAlphaNum,
	&CfEndOfOptions,
	&CfTwoDashWord,
	&CfOneDashLetter,
	&CfOneDashWord,
	&CfWord,
}
