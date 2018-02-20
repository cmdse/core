package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SemanticTokenType model", func() {
	for _, ttype := range SemanticTokenTypes {
		ttype := ttype
		Context(ttype.Name(), func() {
			if ttype.PosModel().Binding == BindRight {
				Context("which is bound right", func() {
					sibling := ttype.Variant().OptValueTokenType()
					It("should have a bound-left option value token type associated with its variant", func() {
						Expect(sibling).ToNot(BeNil())
						Expect(sibling.PosModel().Binding).To(Equal(BindLeft))
					})
					It("should be an option flagName", func() {
						Expect(ttype.PosModel().IsOptionFlag).To(BeTrue())
					})
				})
			}
			if ttype.PosModel().Binding == BindLeft {
				Context("which is bound left", func() {
					sibling := ttype.Variant().FlagTokenType()
					It("should have a bound-right flagName token type associated with its variant", func() {
						Expect(sibling).ToNot(BeNil())
						Expect(sibling.PosModel().Binding).To(Equal(BindRight))
					})
					It("should not be an option flagName", func() {
						Expect(ttype.PosModel().IsOptionFlag).To(BeFalse())
					})
					It("should be an option part", func() {
						Expect(ttype.PosModel().IsOptionPart).To(BeTrue())
					})
				})
			}
			It("should not be bound to Unknown", func() {
				Expect(ttype.PosModel().Binding).NotTo(Equal(BindUnknown))
			})
		})
	}
})
