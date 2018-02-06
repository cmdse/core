package argparse

import (
	"testing"
)

var expectedTokens = map[string]*ContextFreeTokenType{
	"--option=value":      &CF_GNU_EXPLICIT_ASSIGNMENT,
	"--option=12":         &CF_GNU_EXPLICIT_ASSIGNMENT,
	"--long-option=value": &CF_GNU_EXPLICIT_ASSIGNMENT,
	"--long_option=value": &CF_GNU_EXPLICIT_ASSIGNMENT,
	"--long_option=12":    &CF_GNU_EXPLICIT_ASSIGNMENT,
	"--po=TOTO_to":        &CF_GNU_EXPLICIT_ASSIGNMENT,
	"-option=value":       &CF_X2LKT_EXPLICIT_ASSIGNMENT,
	"-option=12":          &CF_X2LKT_EXPLICIT_ASSIGNMENT,
	"-long-option=value":  &CF_X2LKT_EXPLICIT_ASSIGNMENT,
	"-long_option=value":  &CF_X2LKT_EXPLICIT_ASSIGNMENT,
	"+option":             &CF_X2LKT_REVERSE_SWITCH,
	"+long-option":        &CF_X2LKT_REVERSE_SWITCH,
	"+long_option":        &CF_X2LKT_REVERSE_SWITCH,
	"--":                  &CF_END_OF_OPTIONS,
	"-o":                  &ONE_DASH_LETTER,
	"-ns.flag":            &ONE_DASH_WORD, // go cli style namespaced flags
	"-n3":                 &ONE_DASH_WORD, // Typical stick value assignment
	"-n12":                &ONE_DASH_WORD,
	"--option":            &TWO_DASH_WORD,
	"--long-option":       &TWO_DASH_WORD,
	"-_not_an_option":     &WORD,
	"--_not_an_option":    &WORD,
	"word":                &WORD,
	"word with spaces":    &WORD,
}

func TestParseArgument(t *testing.T) {
	for arg, expectedType := range expectedTokens {
		foundType := *ParseArgument(arg)
		if expectedType.Name() != foundType.Name() {
			t.Errorf("ParseArgument didn't convert argument %v to expected type %v but to type %v", arg, expectedType, foundType)
		}
	}
}
