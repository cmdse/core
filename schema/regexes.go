package schema

import (
	"errors"
	"fmt"
	"regexp"
)

type ParametricRegexBuilder struct {
	regexString       string
	defaultFlagMatch  string
	defaultParamMatch string
}

func (parRegex *ParametricRegexBuilder) assembleString(flag string, param string) (string, error) {
	var fullExpr string
	if flag == "" {
		return "", errors.New("flag argument cannot be empty")
	}
	if param == "" {
		fullExpr = fmt.Sprintf(parRegex.regexString, flag)
	} else {
		fullExpr = fmt.Sprintf(parRegex.regexString, flag, param)
	}
	return fullExpr, nil
}

// Build a new regex from parameters flag and param
// If param is empty string, defaults to defaultParamMatch
func (parRegex *ParametricRegexBuilder) Build(flag string, param string) (*regexp.Regexp, error) {
	resParam := param
	if param == "" && parRegex.defaultParamMatch != "" {
		resParam = parRegex.defaultParamMatch
	}
	fullExpr, err := parRegex.assembleString(flag, resParam)
	if err != nil {
		return nil, err
	}
	return regexp.Compile(fullExpr)
}

func (parRegex *ParametricRegexBuilder) BuildDefault() *regexp.Regexp {
	fullExpr, err := parRegex.assembleString(parRegex.defaultFlagMatch, parRegex.defaultParamMatch)
	if err != nil {
		panic(err)
	}
	return regexp.MustCompile(fullExpr)
}

const regexAlphaNumChar = `[A-Za-z0-9]`
const regexAlphaChar = `[A-Za-z]`
const regexNumChar = `[0-9]`
const regexOptionChar = `[\w_\.-]`
const regexValueWordGroup = `.*`

var regexOptionWordGroup = fmt.Sprintf(`%s%s+`, regexAlphaNumChar, regexOptionChar)

var (
	RegBuilderGnuExplicitAssignment = &ParametricRegexBuilder{
		`^--(%s)=(%s)$`,
		regexOptionWordGroup,
		regexValueWordGroup,
	}
	RegBuilderX2lktExplicitAssignment = &ParametricRegexBuilder{
		`^-(%s)=(%s)$`,
		regexOptionWordGroup,
		regexValueWordGroup,
	}
	RegBuilderX2lktReverseSwitch = &ParametricRegexBuilder{
		`^\+(%s)$`,
		regexOptionWordGroup,
		"",
	}
	RegBuilderPosixShortStickyValue = &ParametricRegexBuilder{
		`^-(%s)(%s+)$`,
		regexAlphaChar,
		regexNumChar,
	}
	RegBuilderOneDashLetter = &ParametricRegexBuilder{
		`^-(%s)$`,
		regexAlphaNumChar,
		"",
	}
	RegBuilderOneDashWordAlphaNum = &ParametricRegexBuilder{
		`^-(%s){2,}$`,
		regexAlphaNumChar,
		"",
	}
	RegBuilderOneDashWord = &ParametricRegexBuilder{
		`^-(%s)$`,
		regexOptionWordGroup,
		"",
	}
	RegBuilderTwoDashWord = &ParametricRegexBuilder{
		`^--(%s)$`,
		regexOptionWordGroup,
		"",
	}
	RegBuilderOptWord = &ParametricRegexBuilder{
		`^(%s)$`,
		regexOptionWordGroup,
		"",
	}
	RegBuilderMatchAny = &ParametricRegexBuilder{
		`(%s)`,
		`.*`,
		"",
	}
	RegexGnuExplicitAssignment   = RegBuilderGnuExplicitAssignment.BuildDefault()
	RegexX2lktExplicitAssignment = RegBuilderX2lktExplicitAssignment.BuildDefault()
	RegexX2lktReverseSwitch      = RegBuilderX2lktReverseSwitch.BuildDefault()
	RegexPosixShortStickyValue   = RegBuilderPosixShortStickyValue.BuildDefault()
	RegexOneDashLetter           = RegBuilderOneDashLetter.BuildDefault()
	RegexOneDashWordAlphaNum     = RegBuilderOneDashWordAlphaNum.BuildDefault()
	RegexOneDashWord             = RegBuilderOneDashWord.BuildDefault()
	RegexTwoDashWord             = RegBuilderTwoDashWord.BuildDefault()
	RegexMatchAny                = RegBuilderMatchAny.BuildDefault()
	RegexOptWord                 = RegBuilderOptWord.BuildDefault()
	RegexEndOfOptions            = regexp.MustCompile(`^--$`)
)
