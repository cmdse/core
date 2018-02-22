package argparse

import (
	"fmt"

	. "github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func compareTokenArrays(tokens TokenList, types []TokenType, args []string, semanticCandidatesLen []int) (bool, string) {
	if len(types) != len(tokens) {
		return false, fmt.Sprintf("token list and type list are not of the same length")
	}
	if len(semanticCandidatesLen) != len(tokens) {
		return false, fmt.Sprintf("token list and semantic candidates length list are not of the same length")
	}

	for i, token := range tokens {
		ttype := token.Ttype
		length := len(token.SemanticCandidates)
		expectedLen := semanticCandidatesLen[i]
		if ttype != types[i] {
			detailErrMsg := fmt.Sprintf("expected %T '%s' at position %v/%v for token '%s' but found %T '%s'\n\tToken %v candidates: %v\n", types[i], types[i], i, len(tokens), token.Value, ttype, ttype, i, token.SemanticCandidates)
			fullErr := fmt.Sprintf("\nParse error: %s\n\tArgs     : %v\n\tFound    : %s\n\tExpected : %v\n", detailErrMsg, args, tokens.MapToTypes(), types)
			return false, fullErr
		}
		if length != expectedLen {
			return false, fmt.Sprintf("expected %T '%s' at position %v to have %v semantic candidates but found %v instead\n\tToken %v candidates: %v\n", types[i], types[i], i, expectedLen, length, i, token.SemanticCandidates)
		}
	}
	return true, ""
}

var _ = Describe("ParseArguments method", func() {
	DescribeTable("when provided with no pim",
		func(vararg []string, expected []TokenType, semanticCandidatesLen []int) {
			tokens := ParseArguments(vararg, nil)
			equal, err := compareTokenArrays(tokens, expected, vararg, semanticCandidatesLen)
			Expect(equal).To(BeTrue(), err)
		},
		Entry("should match POSIX and GNU switch + positional operand",
			[]string{"-l", "-p", "--only", "argument"},
			[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortSwitch, SemGNUSwitch, SemOperand},
			[]int{0, 0, 0, 0},
		),
		Entry("should match POSIX, GNU switch, GNU assignment + positional operand",
			[]string{"-l", "--po=TOTO_to", "--only", "argument"},
			[]TokenType{SemPOSIXShortSwitch, SemGNUExplicitAssignment, SemGNUSwitch, SemOperand},
			[]int{0, 0, 0, 0},
		),
		Entry("should handle special non-terminal positional operands",
			[]string{"--po=TOTO_to", "operand", "--only", "argument"},
			[]TokenType{SemGNUExplicitAssignment, SemOperand, SemGNUSwitch, SemOperand},
			[]int{0, 0, 0, 0},
		),
		Entry("should handle end-of-options special switch",
			[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
			[]TokenType{CfOneDashWordAlphaNum, CfOneDashWord, SemEndOfOptions, CfWord, SemOperand, SemOperand},
			[]int{2, 2, 0, 4, 0, 0},
		),
		Entry("should handle end-of-options special switch at last pos",
			[]string{"-option", "-long-option", "--"},
			[]TokenType{CfOneDashWordAlphaNum, SemX2lktSwitch, SemEndOfOptions},
			[]int{2, 0, 0},
		),
	)
	DescribeTable("when provided with program option scheme",
		func(vararg []string, expected []TokenType, scheme OptionScheme, semanticCandidatesLen []int) {
			pim := NewProgramInterfaceModel(scheme, nil)
			tokens := ParseArguments(vararg, pim)
			equal, err := compareTokenArrays(tokens, expected, vararg, semanticCandidatesLen)
			// Expect SemanticCandidates length of context-free tokens > 0
			for _, token := range tokens.WhenContextFree() {
				if _, ok := token.Ttype.(*ContextFreeTokenType); ok {
					Expect(len(token.SemanticCandidates)).To(BeNumerically(">", 0))
				}
			}
			Expect(equal).To(BeTrue(), err)
		},
		Entry("should handle properly when provided with XToolkitStrict option scheme",
			[]string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"},
			[]TokenType{SemX2lktSwitch, CfOneDashWord, SemEndOfOptions, CfWord, SemOperand, SemOperand},
			OptSchemeXToolkitStrict,
			[]int{0, 2, 0, 2, 0, 0},
		),
		Entry("should handle properly when provided with POSIXStrict option scheme",
			[]string{"-xlf", "-p", "optionValue", "-q", "arg1", "arg2"},
			[]TokenType{SemPOSIXStackedShortSwitches, CfOneDashLetter, CfOptWord, CfOneDashLetter, CfOptWord, SemOperand},
			OptionSchemePOSIXStrict,
			[]int{0, 2, 2, 2, 2, 0},
		),
	)
	DescribeTable("when provided with program description model",
		func(vararg []string, expected []TokenType, descriptionModel OptDescriptionModel, semanticCandidatesLen []int) {
			pim := NewProgramInterfaceModel(nil, descriptionModel)
			tokens := ParseArguments(vararg, pim)
			equal, err := compareTokenArrays(tokens, expected, vararg, semanticCandidatesLen)
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
			[]int{0, 0, 0, 0, 0, 0},
		),
		Entry("should bind a left side assignment to the closest token after an end-of-option token",
			[]string{"-x", "-p", "optionValue", "-p", "--", "pArgument", "operand"},
			[]TokenType{SemPOSIXShortSwitch, SemPOSIXShortAssignmentLeftSide, SemPOSIXShortAssignmentValue, SemPOSIXShortAssignmentLeftSide, SemEndOfOptions, SemPOSIXShortAssignmentValue, SemOperand},
			OptDescriptionModel{
				NewOptDescription("execute", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x")),
				NewOptDescription("parse", NewStandaloneMatchModel(VariantPOSIXShortAssignment, "p")),
				NewOptDescription("query", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "q")),
			},
			[]int{0, 0, 0, 0, 0, 0, 0},
		),
	)
})
