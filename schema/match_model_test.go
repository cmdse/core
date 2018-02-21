package schema

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptDescription", func() {
	Describe("NewStandaloneMatchModel function", func() {
		When("given a non-zero flagName", func() {
			model := NewStandaloneMatchModel(VariantX2lktExplicitAssignment, "credentials")
			It("should build leftSideRegex from flagName", func() {
				Expect(model.leftSideRegex).To(BeAssignableToTypeOf(regexp.MustCompile("r")))
			})
			It("should have assigned a variant", func() {
				Expect(model.variant).To(BeAssignableToTypeOf(VariantGNUSwitch))
			})
			It("should have an empty paramName", func() {
				Expect(model.paramName).To(Equal(""))
			})
			It("should have an non-empty flagName", func() {
				Expect(model.flagName).To(Equal("credentials"))
			})
		})
		When("given a zero flagName", func() {
			It("should panic", func() {
				assigningZeroFlag := func() {
					NewStandaloneMatchModel(VariantGNUSwitch, "")
				}
				Expect(assigningZeroFlag).To(Panic())
			})
		})
	})
	newOptDescription := func(models ...*MatchModel) *OptDescription {
		return NewOptDescription("", models...)
	}
	Describe("NewMatchModelFromDefinition function", func() {
		When("given a definition", func() {
			It("should return a corresponding MatchModel", func() {
				matchModel := NewMatchModelFromDefinition(&OptionDefinition{
					VariantX2lktExplicitAssignment,
					"foo",
					"bar",
				})
				Expect(matchModel.variant).To(Equal(VariantX2lktExplicitAssignment))
				Expect(matchModel.flagName).To(Equal("foo"))
				Expect(matchModel.paramName).To(Equal("bar"))
				Expect(matchModel.leftSideRegex.MatchString("-foo=béhéhéhé")).To(BeTrue())
				Expect(matchModel.leftSideRegex.MatchString("-foo=blahblah-blah")).To(BeTrue())
			})
		})
	})
	Describe("NewAssignmentMatchModel function", func() {
		When("given a zero flagName", func() {
			It("should panic", func() {
				assigningZeroFlag := func() {
					NewAssignmentMatchModel(VariantGNUSwitch, "", "blah")
				}
				Expect(assigningZeroFlag).To(Panic())
			})
		})
		When("given a non-zero paramName and non-zero flagName", func() {
			model := NewAssignmentMatchModel(VariantX2lktExplicitAssignment, "credentials", "[a-z]")
			It("should build leftSideRegex from flagName", func() {
				Expect(model.leftSideRegex).To(BeAssignableToTypeOf(regexp.MustCompile("r")))
			})
			It("should have assigned a variant", func() {
				Expect(model.variant).To(BeAssignableToTypeOf(VariantGNUSwitch))
			})
			It("should have a non-empty paramName", func() {
				Expect(model.paramName).To(Equal("[a-z]"))
			})
			It("should have an non-empty flagName", func() {
				Expect(model.flagName).To(Equal("credentials"))
			})
		})
		When("given a non-zero paramName and a zero flagName", func() {
			It("should not panic", func() {
				assigningZeroParam := func() {
					NewAssignmentMatchModel(VariantGNUExplicitAssignment, "credentials", "")
				}
				Expect(assigningZeroParam).ToNot(Panic())
			})
			It("should fallback to default paramName leftSideRegex", func() {
				model := NewAssignmentMatchModel(VariantGNUExplicitAssignment, "credentials", "")
				Expect(model.leftSideRegex.MatchString("--credentials=something")).To(BeTrue())
			})
		})
	})
	Describe("NewPOSIXStackMatchModel function", func() {
		When("given an OptDescriptionModel which have more then two POSIXShortSwitch variants", func() {
			optDscrModel := OptDescriptionModel{
				newOptDescription(NewStandaloneMatchModel(VariantPOSIXShortSwitch, "a")),
				newOptDescription(NewStandaloneMatchModel(VariantPOSIXShortSwitch, "b")),
				newOptDescription(NewStandaloneMatchModel(VariantPOSIXShortSwitch, "c")),
				newOptDescription(NewStandaloneMatchModel(VariantPOSIXShortSwitch, "d")),
			}
			model := NewPOSIXStackMatchModel(optDscrModel)
			DescribeTable("token output",
				func(arg string, shouldMatch bool) {
					Expect(model.MatchLeftSide(arg)).To(Equal(shouldMatch))
				},
				Entry("should match POSIXStacks combinations from the given OptDescriptionModel (1)",
					"-abcd", true,
				),
				Entry("should match POSIXStacks combinations from the given OptDescriptionModel (2)",
					"-ab", true,
				),
				Entry("should match POSIXStacks combinations from the given OptDescriptionModel (3)",
					"-cd", true,
				),
				Entry("should match POSIXStacks combinations from the given OptDescriptionModel (4)",
					"-dbac", true,
				),
				Entry("should not match POSIXStacks combinations that are not in the given OptDescriptionModel (1)",
					"-efgh", false,
				),
				Entry("should not match POSIXStacks combinations that are not in the given OptDescriptionModel (2)",
					"-fg", false,
				),

				Entry("should not match POSIXStacks combinations that are not all in the given OptDescriptionModel (1)",
					"-bhij", false,
				),
				Entry("should not match POSIXStacks combinations that are not all in the given OptDescriptionModel (2)",
					"-def", false,
				),
			)
		})
	})
})
