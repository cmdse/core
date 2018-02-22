package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptDescriptionModel", func() {
	model := OptDescriptionModel{
		NewOptDescription("execute", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x")),
		NewOptDescription("parse", NewStandaloneMatchModel(VariantPOSIXShortAssignment, "p")),
		NewOptDescription("query", NewStandaloneMatchModel(VariantPOSIXShortSwitch, "q")),
	}
	Describe("MatchArgument method", func() {
		It("should match when one of the option description matches", func() {
			candidates1, _ := model.MatchArgument("-x")
			candidates2, _ := model.MatchArgument("-q")
			candidates3, _ := model.MatchArgument("-p")
			Expect(candidates1).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(candidates2).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(candidates3).To(ConsistOf(VariantPOSIXShortAssignment.flagTokenType))
		})
		It("should not match when none of the option description matches", func() {
			Expect(model.MatchArgument("-no-match")).To(HaveLen(0))
		})
	})
	Describe("Variants method", func() {
		It("should return a slice of unique variants", func() {
			Expect(model.Variants()).To(ConsistOf(VariantPOSIXShortSwitch, VariantPOSIXShortAssignment))
			Expect(model.Variants()).To(HaveLen(2))
		})
	})
})
