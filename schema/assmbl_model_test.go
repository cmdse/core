package schema

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newOptParts(variant *OptExpressionVariant, parts ...string) TokenList {
	optionParts := TokenList{}
	flag := NewToken(0, variant.flagTokenType, parts[0], optionParts)
	optionParts = append(optionParts, flag)
	if len(parts) > 1 {
		value := NewToken(1, variant.optValueTokenType, parts[1], optionParts)
		optionParts = append(optionParts, value)
	}
	return optionParts
}

func testOneOpt(model *ExprAssemblyModel, expectedFlag string, expectedValue string, variant *OptExpressionVariant, parts ...string) {
	optionParts := newOptParts(variant, parts...)
	optionExpression, err := model.Assemble(optionParts, variant)
	It("should not return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})
	It("should return exactly 1 option expression", func() {
		Expect(optionExpression.options).To(HaveLen(1), fmt.Sprintf("%v", optionExpression))
	})
	It("its option expression should be composed of flag and value", func() {
		opt := optionExpression.options[0]
		Expect(opt.flag).To(Equal(expectedFlag))
		Expect(opt.assignmentValue).To(Equal(expectedValue))
	})
}

var _ = Describe("ExprAssemblyModel", func() {
	Describe("AssmbModelSplit", func() {
		When("given X-Toolkit explicit assignment regex", func() {
			testOneOpt(AssmbModelSplit, "option", "value", VariantX2lktExplicitAssignment, "-option=value")
		})
		When("given POSIX sticky value assignment regex", func() {
			testOneOpt(AssmbModelSplit, "p", "12", VariantPOSIXShortStickyValue, "-p12")
		})
		When("given a wrong number of arguments", func() {
			optionParts := newOptParts(VariantGNUExplicitAssignment, "--opt=val", "unwantedarg")
			_, err := AssmbModelSplit.Assemble(optionParts, VariantGNUExplicitAssignment)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})
	Describe("AssmbModelGroup", func() {
		When("given GNU implicit assignment regex", func() {
			testOneOpt(AssmbModelGroup, "foo", "bar", VariantGNUImplicitAssignment, "--foo", "bar")
		})
		When("given X-Toolkit implicit assignment regex", func() {
			testOneOpt(AssmbModelGroup, "foo", "bar", VariantX2lktImplicitAssignment, "-foo", "bar")
		})
		When("given a regex matching an unexpected number of groups", func() {
			optionParts := newOptParts(VariantX2lktExplicitAssignment, "--opt=val", "value")
			_, err := AssmbModelSplit.Assemble(optionParts, VariantX2lktExplicitAssignment)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())

			})
		})
	})
	Describe("AssmbModelFlag", func() {
		When("given GNU switch regex", func() {
			testOneOpt(AssmbModelFlag, "foo", "", VariantGNUSwitch, "--foo")
		})
		When("given POSIX switch regex", func() {
			testOneOpt(AssmbModelFlag, "f", "", VariantPOSIXShortSwitch, "-f")
		})
	})
	Describe("AssmbModelFlagStack", func() {
		When("given GNU switch stack regex", func() {
			optionParts := newOptParts(VariantPOSIXStackedShortSwitches, "-pqrst")
			optionExpression, err := AssmbModelFlagStack.Assemble(optionParts, VariantPOSIXStackedShortSwitches)
			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return exactly 5 option expressions", func() {
				Expect(optionExpression.options).To(HaveLen(5))
			})
		})
	})
})
