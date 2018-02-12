package argparse

import (
	. "cmdse-cli/schema"
	"fmt"
	"testing"
)

type simpleTest struct {
	arguments     []string
	expectedTypes []TokenType
}

type testWithScheme struct {
	arguments     []string
	expectedTypes []TokenType
	scheme        OptionScheme
}

type testWithDescrModel struct {
	arguments        []string
	expectedTypes    []TokenType
	descriptionModel OptDescriptionModel
}

var testCore = []simpleTest{
	{
		[]string{"-l", "-p", "--only", "argument"},
		[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortSwitch, SemGNUSwitch, SemOperand},
	}, {
		[]string{"-l", "--po=TOTO_to", "--only", "argument"},
		[]TokenType{SemPOSIXShortSwitch, SemGNUExplicitAssignment, SemGNUSwitch, SemOperand},
	}, {
		[]string{"--po=TOTO_to", "SemOperand", "--only", "argument"},
		[]TokenType{SemGNUExplicitAssignment, SemOperand, SemGNUSwitch, SemOperand},
	}, {
		[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
		[]TokenType{CfOneDashWordAlphaNum, SemX2lktSwitch, SemEndOfOptions, SemOperand, SemOperand, SemOperand},
	},
}

var testWthScheme = []testWithScheme{
	{
		[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
		[]TokenType{SemX2lktSwitch, SemX2lktSwitch, SemEndOfOptions, SemOperand, SemOperand, SemOperand},
		OptSchemeXToolkitStrict,
	},
	{
		[]string{"-xlf", "-p", "optionValue", "-q", "arg1", "arg2"},
		[]TokenType{SemPOSIXStackedShortSwitches, CfOneDashLetter, CfOptWord, CfOneDashLetter, CfOptWord, SemOperand},
		OptionSchemePOSIXStrict,
	},
}

var testWithOpDescr = []testWithDescrModel{
	{
		[]string{"-x", "-p", "optionValue", "-q", "arg1", "arg2"},
		[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortAssignmentLeftSide, SemPOSIXShortAssignmentValue, SemPOSIXShortSwitch, SemOperand, SemOperand},
		OptDescriptionModel{
			&OptDescription{
				"execute",
				[]*MatchModel{
					NewSimpleMatchModel(VariantPOSIXShortSwitch, "x"),
				},
			},
			&OptDescription{
				"parse",
				[]*MatchModel{
					NewSimpleMatchModel(VariantPOSIXShortAssignment, "p"),
				},
			},
			&OptDescription{
				"query",
				[]*MatchModel{
					NewSimpleMatchModel(VariantPOSIXShortSwitch, "q"),
				},
			},
		},
	},
}

func compareTokenArrays(tokens TokenList, types []TokenType) (isEqual bool, err error) {
	if len(types) != len(tokens) {
		return false, fmt.Errorf("token list and type list are not of the same length")
	}
	for i, token := range tokens {
		ttype := token.ttype
		if ttype != types[i] {
			return false, fmt.Errorf("expected %T '%s' at position %v for token '%s' but found %T '%s'\n\tToken %v candidates: %v", types[i], types[i], i, token.value, ttype, ttype, i, token.semanticCandidates)
		}
	}
	return true, nil
}

func TestParseArguments(t *testing.T) {

	for i, test := range testCore {
		tokens := ParseArguments(test.arguments, nil)
		equal, err := compareTokenArrays(tokens, test.expectedTypes)
		if !equal {
			t.Errorf("Parsing test #%v error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v", i, err, test.arguments, tokens.MapToTypes(), test.expectedTypes)
		}

	}
}

func TestParseWithOptionScheme(t *testing.T) {

	for i, test := range testWthScheme {
		scheme := test.scheme
		pim := NewProgramInterfaceModel(scheme, nil)
		tokens := ParseArguments(test.arguments, pim)
		equal, err := compareTokenArrays(tokens, test.expectedTypes)
		if !equal {
			t.Errorf("Parsing test #%v error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v", i, err, test.arguments, tokens.MapToTypes(), test.expectedTypes)
		}

	}
}

func TestParseWithOptionDescriptionModel(t *testing.T) {

	for i, test := range testWithOpDescr {
		descriptionModel := test.descriptionModel
		pim := NewProgramInterfaceModel(nil, descriptionModel)
		tokens := ParseArguments(test.arguments, pim)
		equal, err := compareTokenArrays(tokens, test.expectedTypes)
		if !equal {
			t.Errorf("Parsing test #%v error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v", i, err, test.arguments, tokens.MapToTypes(), test.expectedTypes)
		}

	}
}
