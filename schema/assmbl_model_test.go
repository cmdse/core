package schema

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type optionParts struct {
	args []string
}

func (op *optionParts) Args() []string {
	return op.args
}

func newOptionParts(parts ...string) OptionParts {
	return &optionParts{
		parts,
	}
}

func checkOneOpt(model *ExprAssemblyModel, expectedFlag string, expectedValue string, reg *regexp.Regexp, parts ...string) {
	optionParts := newOptionParts(parts...)
	optionExpression, err := model.Assemble(optionParts, reg)
	It("should not return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})
	It("should return exactly 1 option expression", func() {
		Expect(optionExpression.options).To(HaveLen(1))
	})
	It("the option expression should be composed of flag and value", func() {
		opt := optionExpression.options[0]
		Expect(opt.flag).To(Equal(expectedFlag))
		Expect(opt.assignmentValue).To(Equal(expectedValue))
	})
}

var _ = Describe("ExprAssemblyModel", func() {
	Describe("AssmbModelSplit", func() {
		When("given X-Toolkit explicit assignment regex", func() {
			checkOneOpt(AssmbModelSplit, "option", "value", RegexX2lktExplicitAssignment, "-option=value")
		})
		When("given POSIX sticky value assignment regex", func() {
			checkOneOpt(AssmbModelSplit, "p", "12", RegexPosixShortStickyValue, "-p12")
		})
		When("given a wrong number of arguments", func() {
			optionParts := newOptionParts("--opt=val", "unwantedarg")
			_, err := AssmbModelSplit.Assemble(optionParts, RegexX2lktExplicitAssignment)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})
	Describe("AssmbModelGroup", func() {
		When("given GNU implicit assignment regex", func() {
			checkOneOpt(AssmbModelGroup, "foo", "bar", RegexTwoDashWord, "--foo", "bar")
		})
		When("given X-Toolkit implicit assignment regex", func() {
			checkOneOpt(AssmbModelGroup, "foo", "bar", RegexOneDashWord, "-foo", "bar")
		})
		When("given a regex matching an unexpected number of groups", func() {
			optionParts := newOptionParts("--opt=val", "value")
			_, err := AssmbModelSplit.Assemble(optionParts, RegexX2lktExplicitAssignment)
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())

			})
		})
	})
	Describe("AssmbModelFlag", func() {
		When("given GNU switch regex", func() {
			checkOneOpt(AssmbModelFlag, "foo", "", RegexOneDashWord, "-foo")
		})
		When("given POSIX switch regex", func() {
			checkOneOpt(AssmbModelFlag, "f", "", RegexOneDashLetter, "-f")
		})
	})
	Describe("AssmbModelFlagStack", func() {
		When("given GNU switch stack regex", func() {
			optionParts := newOptionParts("-pqrst")
			optionExpression, err := AssmbModelFlagStack.Assemble(optionParts, RegexOneDashWordAlphaNum)
			It("should not return an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should return exactly 5 option expressions", func() {
				Expect(optionExpression.options).To(HaveLen(5))
			})
		})
	})
})
