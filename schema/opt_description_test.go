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
				matchedCandidates, matched := description.MatchArgument("--execute")
				Expect(matched).To(BeTrue())
				Expect(matchedCandidates).To(ConsistOf(SemGNUSwitch))
			})
			It("should match x-toolkit switch", func() {
				matchedCandidates, matched := description.MatchArgument("-execute")
				Expect(matched).To(BeTrue())
				Expect(matchedCandidates).To(ConsistOf(SemX2lktSwitch))
			})
			It("should match POSIX switch", func() {
				matchedCandidates, matched := description.MatchArgument("-x")
				Expect(matched).To(BeTrue())
				Expect(matchedCandidates).To(ConsistOf(SemPOSIXShortSwitch))
			})
			It("should not match when no match model exists", func() {
				matchedCandidates, matched := description.MatchArgument("-p")
				Expect(matched).To(BeFalse())
				Expect(matchedCandidates).To(HaveLen(0))
			})
		})
		When("provided with match models composed of switches and options assignments", func() {
			description := NewOptDescription("execute",
				NewStandaloneMatchModel(VariantPOSIXShortAssignment, "x"),
				NewStandaloneMatchModel(VariantPOSIXShortSwitch, "x"),
			)
			It("should match multiple token types when they should both match", func() {
				matchedCandidates, matched := description.MatchArgument("-x")
				Expect(matched).To(BeTrue())
				Expect(matchedCandidates).To(ConsistOf(SemPOSIXShortAssignmentLeftSide, SemPOSIXShortSwitch))
			})
		})
	})
})
