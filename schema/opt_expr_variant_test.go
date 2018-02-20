package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptExpressionVariant", func() {
	Describe("Build method", func() {
		It("should panic when variant has an assembly model type 'Flag' and a paramName is provided", func() {
			Expect(func() { VariantGNUSwitch.Build("something", []string{"toto"}) }).To(Panic())
		})
		It("should not panic when provided flagName is an invalid regex string (should be quoted)", func() {
			Expect(func() { VariantX2lktExplicitAssignment.Build("(((((", []string{"toto"}) }).ToNot(Panic())
		})
		It("should return a regex matching the variant", func() {
			regex := VariantX2lktExplicitAssignment.Build("exec", nil)
			Expect(regex.MatchString("-exec=/exec/path")).To(BeTrue())
		})
	})
})
