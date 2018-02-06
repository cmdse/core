package argparse

import (
	"fmt"
	"testing"
)

type testTuple struct {
	arguments     []string
	expectedTypes []TokenType
}

var tests = []testTuple{
	{
		[]string{"-l", "-p", "--only", "argument"},
		[]TokenType{POSIX_SHORT_SWITCH, POSIX_SHORT_SWITCH, GNU_SWITCH, OPERAND},
	}, {
		[]string{"-l", "--po=TOTO_to", "--only", "argument"},
		[]TokenType{POSIX_SHORT_SWITCH, GNU_EXPLICIT_ASSIGNMENT, GNU_SWITCH, OPERAND},
	}, {
		[]string{"--po=TOTO_to", "operand", "--only", "argument"},
		[]TokenType{GNU_EXPLICIT_ASSIGNMENT, OPERAND, GNU_SWITCH, OPERAND},
	},
}

func compareTokenArrays(tokens TokenList, types []TokenType) (isEqual bool, err error) {
	if len(types) != len(tokens) {
		return false, fmt.Errorf("token list and type list are not of the same length")
	}
	for i, token := range tokens {
		ttype := *token.ttype
		if ttype != types[i] {
			return false, fmt.Errorf("expected %T '%s' at position %v for token '%s' but found %T '%s'", types[i], types[i], i, token.value, ttype, ttype)
		}
	}
	return true, nil
}

func TestParseArguments(t *testing.T) {

	for i, test := range tests {
		tokens := ParseArguments(test.arguments)
		equal, err := compareTokenArrays(tokens, test.expectedTypes)
		if !equal {
			t.Errorf("Parsing test #%v error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v", i, err, test.arguments, tokens.MapToTypes(), test.expectedTypes)
		}

	}
}
