package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
)

var _ = Describe("OptDescription", func() {
	Describe("NewSimpleMatchModel function", func() {
		model := NewSimpleMatchModel(VariantX2lktExplicitAssignment, "credentials")
		It("should build regex from flag", func() {
			Expect(model.regex).To(BeAssignableToTypeOf(regexp.MustCompile("r")))
		})
		It("should have assigned a variant", func() {
			Expect(model.variant).To(BeAssignableToTypeOf(VariantGNUSwitch))
		})
		It("should have an empty param", func() {
			Expect(model.param).To(Equal(""))
		})
		It("should have an non-empty flag", func() {
			Expect(model.flag).To(Equal("credentials"))
		})
	})
	Describe("NewMatchModelWithTypedValue function", func() {
		model := NewMatchModelWithTypedValue(VariantX2lktExplicitAssignment, "credentials", "[a-z]")
		It("should build regex from flag", func() {
			Expect(model.regex).To(BeAssignableToTypeOf(regexp.MustCompile("r")))
		})
		It("should have assigned a variant", func() {
			Expect(model.variant).To(BeAssignableToTypeOf(VariantGNUSwitch))
		})
		It("should have a non-empty param", func() {
			Expect(model.param).To(Equal("[a-z]"))
		})
		It("should have an non-empty flag", func() {
			Expect(model.flag).To(Equal("credentials"))
		})
	})
})
