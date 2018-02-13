package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//func TestSemanticModelConsistency(t *testing.T) {
//	for _, ttype := range SemanticTokenTypes {
//		if ttype.PosModel().Binding == BindRight {
//			sibling := ttype.Variant().OptValueTokenType()
//			if sibling == nil {
//				t.Errorf("%v : a bound-right token type must have an opt value token type associated with its variant.", ttype.Name())
//			}
//			if !ttype.PosModel().IsOptionFlag {
//				t.Errorf("%v : a bound-right token type should be an option flag. ", ttype.Name())
//			}
//		}
//		if ttype.PosModel().Binding == BindLeft {
//			sibling := ttype.Variant().FlagTokenType()
//			if sibling == nil {
//				t.Errorf("%v : a bound-left token type must have a flag token type associated with its variant.", ttype.Name())
//			}
//			if ttype.PosModel().IsOptionFlag {
//				t.Errorf("%v : a bound-left token type cannot be an option flag.", ttype.Name())
//			}
//			if !ttype.PosModel().IsOptionPart {
//				t.Errorf("%v : a bound-left token type should be an option flag.", ttype.Name())
//			}
//		}
//		if ttype.PosModel().Binding == BindUnknown {
//			t.Errorf("%v : a semantic token type cannot have an unknown binding.", ttype.Name())
//		}
//	}
//}

var _ = Describe("SemanticTokenType", func() {
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
					It("should be an option flag", func() {
						Expect(ttype.PosModel().IsOptionFlag).To(BeTrue())
					})
				})
			}
			if ttype.PosModel().Binding == BindLeft {
				Context("which is bound left", func() {
					sibling := ttype.Variant().FlagTokenType()
					It("should have a bound-right flag token type associated with its variant", func() {
						Expect(sibling).ToNot(BeNil())
						Expect(sibling.PosModel().Binding).To(Equal(BindRight))
					})
					It("should not be an option flag", func() {
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
