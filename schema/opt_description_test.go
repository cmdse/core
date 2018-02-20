package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptDescription", func() {
	Describe("MatchArgument method", func() {
		When("provided with match models composed of switches", func() {
			description := &OptDescription{
				"execute",
				MatchModels{
					NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x"),
					NewStandaloneMatchModel(VariantGNUSwitch, "execute"),
					NewStandaloneMatchModel(VariantX2lktSwitch, "execute"),
				},
			}
			It("should match GNU switch", func() {
				Expect(description.MatchArgument("--execute")).To(ConsistOf(SemGNUSwitch))
			})
			It("should match x-toolkit switch", func() {
				Expect(description.MatchArgument("-execute")).To(ConsistOf(SemX2lktSwitch))
			})
			It("should match POSIX switch", func() {
				Expect(description.MatchArgument("-x")).To(ConsistOf(SemPOSIXShortSwitch))
			})
			It("should not match when no match model exists", func() {
				Expect(description.MatchArgument("-e")).To(BeNil())
			})
		})
		When("provided with match models composed of switches and options assignments", func() {
			description := NewOptDescription("execute",
				NewStandaloneMatchModel(VariantPOSIXShortAssignment, "x"),
				NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x"),
			)
			It("should match multiple token types when they should both match", func() {
				Expect(description.MatchArgument("-x")).To(ConsistOf(SemPOSIXShortAssignmentLeftSide, SemPOSIXShortSwitch))
			})
		})
	})
})
