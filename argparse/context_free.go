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

func (tokenType ContextFreeTokenType) IsSemantic() bool {
	return tokenType.PosModel().IsSemantic
}

func (tokenType ContextFreeTokenType) PosModel() *PositionalModel {
	return &UNSET
}

func (tokenType ContextFreeTokenType) Name() string {
	return tokenType.name
}

func (tokenType ContextFreeTokenType) String() string {
	return tokenType.name
}

func (tokenType ContextFreeTokenType) Equal(comparedTtype *ContextFreeTokenType) bool {
	return tokenType.name == comparedTtype.name
}

func (tokenType ContextFreeTokenType) Regexp() (*regexp.Regexp, bool) {
	// each regex is tested, so compilable regex can be asserted
	reg := regexp.MustCompile(tokenType.regexp)
	isMatchAllRegex := true
	isMatchAllRegex = reg.String() == ""
	return reg, isMatchAllRegex
}

const OptionWordGroup = `([A-Za-z0-9][\w_\.-]+)`
const ValueWordGroup = `(.*)`

var (
	CF_GNU_EXPLICIT_ASSIGNMENT = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&GNU_EXPLICIT_ASSIGNMENT,
		},
		regexp: fmt.Sprintf(`^--%s=%s$`, OptionWordGroup, ValueWordGroup),
		name:   "CF_GNU_EXPLICIT_ASSIGNMENT",
	}
	CF_X2LKT_EXPLICIT_ASSIGNMENT = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&X2LKT_EXPLICIT_ASSIGNMENT,
		},
		regexp: fmt.Sprintf(`^-%s=%s$`, OptionWordGroup, ValueWordGroup),
		name:   "CF_X2LKT_EXPLICIT_ASSIGNMENT",
	}
	CF_X2LKT_REVERSE_SWITCH = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&X2LKT_REVERSE_SWITCH,
		},
		regexp: fmt.Sprintf(`^\+%s$`, OptionWordGroup),
		name:   "CF_X2LKT_REVERSE_SWITCH",
	}
	CF_END_OF_OPTIONS = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&END_OF_OPTIONS,
		},
		regexp: `^--$`,
		name:   "CF_END_OF_OPTIONS",
	}
	ONE_DASH_LETTER = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&POSIX_SHORT_ASSIGNMENT_LEFT_SIDE,
			&POSIX_SHORT_SWITCH,
		},
		regexp: `^-(\w)$`,
		name:   "ONE_DASH_LETTER",
	}
	ONE_DASH_WORD = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&POSIX_STACKED_SHORT_SWITCHES,
			&POSIX_SHORT_STICKY_VALUE,
			&X2LKT_SWITCH,
			&X2LKT_IMPLICIT_ASSIGNEMNT_LEFT_SIDE,
		},
		regexp: fmt.Sprintf(`^-%s$`, OptionWordGroup),
		name:   "ONE_DASH_WORD",
	}
	TWO_DASH_WORD = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&GNU_SWITCH,
			&GNU_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		},
		regexp: fmt.Sprintf(`^--%s$`, OptionWordGroup),
		name:   "TWO_DASH_WORD",
	}
	WORD = ContextFreeTokenType{
		SemanticCandidates: []*SemanticTokenType{
			&OPERAND,
			&POSIX_SHORT_ASSIGNMENT_VALUE,
			&GNU_IMPLICIT_ASSIGNMENT_VALUE,
			&X2LKT_IMPLICIT_ASSIGNMENT_VALUE,
			&HEADLESS_OPTION,
		},
		name: "WORD",
	}
)

var ContextFreeTokenTypes = []*ContextFreeTokenType{
	&CF_GNU_EXPLICIT_ASSIGNMENT,
	&CF_X2LKT_REVERSE_SWITCH,
	&CF_X2LKT_EXPLICIT_ASSIGNMENT,
	&CF_END_OF_OPTIONS,
	&TWO_DASH_WORD,
	&ONE_DASH_LETTER,
	&ONE_DASH_WORD,
	&WORD,
}
