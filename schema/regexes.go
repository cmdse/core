package schema

import (
	"fmt"
	"regexp"
)

const regexAlphaNumChar = `[A-Za-z0-9]`
const regexAlphaChar = `[A-Za-z]`
const regexNumChar = `[0-9]`
const regexOptionChar = `[\w_\.-]`
const regexValueWordGroup = `.*`

var regexOptionWordGroup = fmt.Sprintf(`%s%s+`, regexAlphaNumChar, regexOptionChar)

var (
	RegexGnuExplicitAssignment   = regexp.MustCompile(fmt.Sprintf(`^--(%s)=(%s)$`, regexOptionWordGroup, regexValueWordGroup))
	RegexX2lktExplicitAssignment = regexp.MustCompile(fmt.Sprintf(`^-(%s)=(%s)$`, regexOptionWordGroup, regexValueWordGroup))
	RegexX2lktReverseSwitch      = regexp.MustCompile(fmt.Sprintf(`^\+(%s)$`, regexOptionWordGroup))
	RegexEndOfOptions            = regexp.MustCompile(`^--$`)
	RegexOneDashLetter           = regexp.MustCompile(fmt.Sprintf(`^-(%s)$`, regexAlphaNumChar))
	RegexPosixShortStickyValue   = regexp.MustCompile(fmt.Sprintf(`^-(%s)(%s+)$`, regexAlphaChar, regexNumChar))
	RegexOneDashWordAlphaNum     = regexp.MustCompile(fmt.Sprintf(`^-(%s){2,}$`, regexAlphaNumChar))
	RegexOneDashWord             = regexp.MustCompile(fmt.Sprintf(`^-(%s)$`, regexOptionWordGroup))
	RegexTwoDashWord             = regexp.MustCompile(fmt.Sprintf(`^--(%s)$`, regexOptionWordGroup))
	RegexMatchAny                = regexp.MustCompile(`.*`)
)
