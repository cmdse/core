package argparse

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

const OptionWordGroup = `([A-Za-z0-9][\w_\.-]+)`
const ValueWordGroup = `(.*)`

var (
	CfGnuExplicitAssignment = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemGnuExplicitAssignment,
		},
		regexp: fmt.Sprintf(`^--%s=%s$`, OptionWordGroup, ValueWordGroup),
		name:   "CfGnuExplicitAssignment",
	}
	CfX2lktExplicitAssignment = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemX2lktExplicitAssignment,
		},
		regexp: fmt.Sprintf(`^-%s=%s$`, OptionWordGroup, ValueWordGroup),
		name:   "CfX2lktExplicitAssignment",
	}
	CfX2lktReverseSwitch = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemX2lktReverseSwitch,
		},
		regexp: fmt.Sprintf(`^\+%s$`, OptionWordGroup),
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
		regexp: `^-(\w)$`,
		name:   "CfOneDashLetter",
	}
	CfOneDashWord = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemPosixStackedShortSwitches,
			&SemPosixShortStickyValue,
			&SemX2lktSwitch,
			&SemX2lktImplicitAssignmentLeftSide,
		},
		regexp: fmt.Sprintf(`^-%s$`, OptionWordGroup),
		name:   "CfOneDashWord",
	}
	CfTwoDashWord = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&SemGnuSwitch,
			&SemGnuImplicitAssignmentLeftSide,
		},
		regexp: fmt.Sprintf(`^--%s$`, OptionWordGroup),
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
	&CfEndOfOptions,
	&CfTwoDashWord,
	&CfOneDashLetter,
	&CfOneDashWord,
	&CfWord,
}
