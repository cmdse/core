package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptDescriptionModel", func() {
	model := OptDescriptionModel{
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
	}
	Describe("MatchArgument method", func() {
		It("it should match when one of the option description matches", func() {
			Expect(model.MatchArgument("-x")).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(model.MatchArgument("-q")).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(model.MatchArgument("-p")).To(ConsistOf(VariantPOSIXShortAssignment.flagTokenType))
		})
		It("it should not match when none of the option description matches", func() {
			Expect(model.MatchArgument("-no-match")).To(HaveLen(0))
		})
	})
})
