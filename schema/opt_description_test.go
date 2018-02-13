package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptDescription", func() {
	Describe("Match method", func() {
		When("provided with match models composed of switches", func() {
			description := &OptDescription{
				"execute",
				[]*MatchModel{
					NewSimpleMatchModel(VariantPOSIXShortSwitch, "x"),
					NewSimpleMatchModel(VariantGNUSwitch, "execute"),
					NewSimpleMatchModel(VariantX2lktSwitch, "execute"),
				},
			}
			It("should match GNU switch", func() {
				Expect(description.Match("--execute")).To(ConsistOf(SemGNUSwitch))
			})
			It("should match x-toolkit switch", func() {
				Expect(description.Match("-execute")).To(ConsistOf(SemX2lktSwitch))
			})
			It("should match POSIX switch", func() {
				Expect(description.Match("-x")).To(ConsistOf(SemPOSIXShortSwitch))
			})
			It("should not match when no match model exists", func() {
				Expect(description.Match("-e")).To(BeNil())
			})
		})
		When("provided with match models composed of switches and options assignments", func() {
			description := &OptDescription{
				"execute",
				[]*MatchModel{
					NewSimpleMatchModel(VariantPOSIXShortAssignment, "x"),
					NewSimpleMatchModel(VariantPOSIXShortSwitch, "x"),
				},
			}
			It("should match multiple token types when they should both match", func() {
				Expect(description.Match("-x")).To(ConsistOf(SemPOSIXShortAssignmentLeftSide, SemPOSIXShortSwitch))
			})
		})
	})
})
