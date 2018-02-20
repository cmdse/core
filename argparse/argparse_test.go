package argparse

import (
	"fmt"

	. "github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func compareTokenArrays(tokens TokenList, types []TokenType, args []string) (bool, string) {
	if len(types) != len(tokens) {
		return false, fmt.Sprintf("token list and type list are not of the same length")
	}
	for i, token := range tokens {
		ttype := token.ttype
		if ttype != types[i] {
			detailErrMsg := fmt.Sprintf("expected %T '%s' at position %v/%v for token '%s' but found %T '%s'\n\tToken %v candidates: %v\n", types[i], types[i], i, len(tokens), token.value, ttype, ttype, i, token.semanticCandidates)
			fullErr := fmt.Sprintf("\nParse error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v\n", detailErrMsg, args, tokens.MapToTypes(), types)
			return false, fullErr
		}
	}
	return true, ""
}

var _ = Describe("ParseArguments method", func() {
	When("provided with no pim", func() {
		DescribeTable("token output",
			func(vararg []string, expected []TokenType) {
				tokens := ParseArguments(vararg, nil)
				equal, err := compareTokenArrays(tokens, expected, vararg)
				Expect(equal).To(BeTrue(), err)
			},
			Entry("should match POSIX and GNU switch + positional operand",
				[]string{"-l", "-p", "--only", "argument"},
				[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortSwitch, SemGNUSwitch, SemOperand},
			),
			Entry("should match POSIX, GNU switch, GNU assignment + positional operand",
				[]string{"-l", "--po=TOTO_to", "--only", "argument"},
				[]TokenType{SemPOSIXShortSwitch, SemGNUExplicitAssignment, SemGNUSwitch, SemOperand},
			),
			Entry("",
				[]string{"--po=TOTO_to", "SemOperand", "--only", "argument"},
				[]TokenType{SemGNUExplicitAssignment, SemOperand, SemGNUSwitch, SemOperand},
			),
			Entry("should handle end-of-options special switch",
				[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
				[]TokenType{CfOneDashWordAlphaNum, CfOneDashWord, SemEndOfOptions, CfOneDashWordAlphaNum, SemOperand, SemOperand},
			),
		)
	})
	When("provided with program option scheme", func() {
		DescribeTable("token output",
			func(vararg []string, expected []TokenType, scheme OptionScheme) {
				pim := NewProgramInterfaceModel(scheme, nil)
				tokens := ParseArguments(vararg, pim)
				equal, err := compareTokenArrays(tokens, expected, vararg)
				Expect(equal).To(BeTrue(), err)
			},
			Entry("should handle properly when provided with XToolkitStrict option scheme",
				[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
				[]TokenType{SemX2lktSwitch, CfOneDashWord, SemEndOfOptions, CfOneDashWordAlphaNum, SemOperand, SemOperand},
				OptSchemeXToolkitStrict,
			),
			Entry("should handle properly when provided with POSIXStrict option scheme",
				[]string{"-xlf", "-p", "optionValue", "-q", "arg1", "arg2"},
				[]TokenType{SemPOSIXStackedShortSwitches, CfOneDashLetter, CfOptWord, CfOneDashLetter, CfOptWord, SemOperand},
				OptionSchemePOSIXStrict,
			),
		)
	})
	When("provided with program description model", func() {
		DescribeTable("token output",
			func(vararg []string, expected []TokenType, descriptionModel OptDescriptionModel) {
				pim := NewProgramInterfaceModel(nil, descriptionModel)
				tokens := ParseArguments(vararg, pim)
				equal, err := compareTokenArrays(tokens, expected, vararg)
				Expect(equal).To(BeTrue(), err)
			},
			Entry("should handle properly when provided with a description model matching short switches and assignments",
				[]string{"-x", "-p", "optionValue", "-q", "arg1", "arg2"},
				[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortAssignmentLeftSide, SemPOSIXShortAssignmentValue, SemPOSIXShortSwitch, SemOperand, SemOperand},
				OptDescriptionModel{
					NewOptDescription("execute", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x")),
					NewOptDescription("parse", NewStandaloneMatchModel(VariantPOSIXShortAssignment, "p")),
					NewOptDescription("query", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "q")),
				},
			),
			Entry("should bind a left side assignment to the closest token after an end-of-option token",
				[]string{"-x", "-p", "optionValue", "-p", "--", "pArgument", "operand"},
				[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortAssignmentLeftSide, SemPOSIXShortAssignmentValue, SemPOSIXShortAssignmentLeftSide, SemEndOfOptions, SemPOSIXShortAssignmentValue, SemOperand},
				OptDescriptionModel{
					NewOptDescription("execute", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x")),
					NewOptDescription("parse", NewStandaloneMatchModel(VariantPOSIXShortAssignment, "p")),
					NewOptDescription("query", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "q")),
				},
			),
		)
	})
})
