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
			Expect(model.MatchArgument("-x")).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(model.MatchArgument("-q")).To(ConsistOf(VariantPOSIXShortSwitch.flagTokenType))
			Expect(model.MatchArgument("-p")).To(ConsistOf(VariantPOSIXShortAssignment.flagTokenType))
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
