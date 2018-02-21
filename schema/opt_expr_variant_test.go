package schema

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func testOneOptVariant(expectedFlag string, expectedValue string, variant *OptExpressionVariant, parts ...string) {
	optionParts := newOptParts(variant, parts...)
	optionExpression, err := variant.Assemble(optionParts)
	Expect(err).NotTo(HaveOccurred())
	Expect(optionExpression.options).To(HaveLen(1), fmt.Sprintf("%v", optionExpression))
	opt := optionExpression.options[0]
	Expect(opt.flag).To(Equal(expectedFlag))
	Expect(opt.assignmentValue).To(Equal(expectedValue))
}

var _ = Describe("OptExpressionVariant", func() {
	Describe("Build method", func() {
		It("should panic when variant has an assembly model type 'flag' and a paramName is provided", func() {
			Expect(func() { VariantGNUSwitch.Build("something", []string{"toto"}) }).To(Panic())
		})
		It("should not panic when provided flagName is an invalid regex string (should be quoted)", func() {
			Expect(func() { VariantX2lktExplicitAssignment.Build("(((((", []string{"toto"}) }).ToNot(Panic())
		})
		It("should return a regex matching the variant", func() {
			regex := VariantX2lktExplicitAssignment.Build("exec", nil)
			Expect(regex.MatchString("-exec=/exec/path")).To(BeTrue())
		})
		DescribeTable("Assemble method", func(variant *OptExpressionVariant, expectedFlag string, expectedValue string, parts ...string) {
			testOneOptVariant(expectedFlag, expectedValue, variant, parts...)
		},
			Entry(VariantPOSIXShortSwitch.name, VariantPOSIXShortSwitch,
				"p", "", "-p",
			),
			Entry(VariantPOSIXShortAssignment.name, VariantPOSIXShortAssignment,
				"p", "value", "-p", "value",
			),
			Entry(VariantPOSIXShortStickyValue.name, VariantPOSIXShortStickyValue,
				"p", "12", "-p12",
			),
			Entry(VariantX2lktSwitch.name, VariantX2lktSwitch,
				"switch", "", "-switch",
			),
			Entry(VariantX2lktReverseSwitch.name, VariantX2lktReverseSwitch,
				"switch", "", "+switch",
			),
			Entry(VariantX2lktImplicitAssignment.name, VariantX2lktImplicitAssignment,
				"switch", "value", "-switch", "value",
			),
			Entry(VariantX2lktExplicitAssignment.name, VariantX2lktExplicitAssignment,
				"option", "value", "-option=value",
			),
			Entry(VariantGNUSwitch.name, VariantGNUSwitch,
				"switch", "", "--switch",
			),
			Entry(VariantGNUImplicitAssignment.name, VariantGNUImplicitAssignment,
				"option", "value", "--option", "value",
			),
			Entry(VariantGNUExplicitAssignment.name, VariantGNUExplicitAssignment,
				"option", "value", "--option=value",
			),
			Entry(VariantHeadlessOption.name, VariantHeadlessOption,
				"option", "", "option",
			),
			Entry(VariantEndOfOptions.name, VariantEndOfOptions,
				"", "", "--",
			),
			// VariantPOSIXStackedShortSwitches cannot be tested because it returns multiple option parts
		)
	})
})
