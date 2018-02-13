package argparse

import (
	. "github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token", func() {
	Describe("IsBoundToOneOf func", func() {
		var tokens = TokenList{}
		var tokenType TokenType = CfOneDashWord
		token := NewToken(0, tokenType, "-option", tokens)
		tokens = append(tokens, token)
		When("when token type is context-free", func() {
			Context("and its semantic candidates are bound 'None' and 'Right'", func() {
				Context("and provided bindings are 'None' and 'Right'", func() {
					boundNoneOrRight := token.IsBoundToOneOf(Bindings{BindNone, BindRight})
					It("should return true", func() {
						Expect(boundNoneOrRight).To(BeTrue())
					})
				})
				Context("and provided binding are 'None' and 'Left'", func() {
					boundNoneOrLeft := token.IsBoundToOneOf(Bindings{BindNone, BindLeft})
					It("should return false", func() {
						Expect(boundNoneOrLeft).To(BeFalse())
					})
				})
				Context("and provided bindings are 'Unknown' and 'Left'", func() {
					boundUnknownOrLeft := token.IsBoundToOneOf(Bindings{BindUnknown, BindLeft})
					It("should return false", func() {
						Expect(boundUnknownOrLeft).To(BeFalse())
					})
				})
			})
		})
		When("when token type is semantic and bound 'None'", func() {
			var tokens = TokenList{}
			var tokenType TokenType = SemX2lktSwitch
			token := NewToken(0, tokenType, "-option", tokens)
			tokens = append(tokens, token)
			Context("and provided bindings are 'None' and 'Left'", func() {
				boundToNoneOrLeft := token.IsBoundToOneOf(Bindings{BindNone, BindLeft})
				It("should return true", func() {
					Expect(boundToNoneOrLeft).To(BeTrue())
				})
			})
			Context("and provided bindings are 'Unknown' and 'Right'", func() {
				boundToUnknownOrRight := token.IsBoundToOneOf(Bindings{BindUnknown, BindRight})
				It("should return false", func() {
					Expect(boundToUnknownOrRight).To(BeFalse())
				})
			})
		})

	})
	Describe("IsBoundTo func", func() {
		var tokens = TokenList{}
		var tokenType TokenType = CfEndOfOptions
		token := NewToken(0, tokenType, "--", tokens)
		tokens = append(tokens, token)
		Context("when token type is semantic and bound to 'none'", func() {
			When("and provided binding is 'None'", func() {
				It("should return true", func() {
					boundToNone := token.IsBoundTo(BindNone)
					Expect(boundToNone).To(BeTrue())
				})
			})
			When("and provided binding is 'Left'", func() {
				It("should return false", func() {
					boundToLeft := token.IsBoundTo(BindLeft)
					Expect(boundToLeft).To(BeFalse())
				})
			})
		})
		Context("when token type is context-free and its semantic candidates are bound to 'None'", func() {
			var tokens = TokenList{}
			var tokenType TokenType = CfEndOfOptions
			token := NewToken(0, tokenType, "-test", tokens)
			tokens = append(tokens, token)
			When("when provided with 'None' binding", func() {
				It("should return true", func() {
					boundToNone := token.IsBoundTo(BindNone)
					Expect(boundToNone).To(BeTrue())
				})
			})
			When("when provided with 'Left' binding", func() {
				It("should return false", func() {
					boundToLeft := token.IsBoundTo(BindLeft)
					Expect(boundToLeft).To(BeFalse())
				})
			})

		})
	})
	Describe("IsOptionPart func", func() {
		When("when token is context free and its semantic candidates are option parts", func() {
			var tokens = TokenList{}
			var tokenType TokenType = CfOneDashWord
			token := NewToken(0, tokenType, "-test", tokens)
			tokens = append(tokens, token)
			isOption := token.IsOptionPart()
			It("should return true", func() {
				Expect(isOption).To(BeTrue())
			})
		})
		When("when token is semantic and option part", func() {
			var tokens = TokenList{}
			var tokenType TokenType = SemX2lktExplicitAssignment
			token := NewToken(0, tokenType, "-opt=value", tokens)
			tokens = append(tokens, token)
			isOptionPart := token.IsOptionPart()
			It("should return true", func() {
				Expect(isOptionPart).To(BeTrue())
			})
		})
	})
	Describe("IsOptionFlag func", func() {
		When("when token is context free and its semantic candidates are option flags", func() {
			var tokens = TokenList{}
			var tokenType TokenType = CfEndOfOptions
			token := NewToken(0, tokenType, "--", tokens)
			tokens = append(tokens, token)
			isOption := token.IsOptionFlag()
			It("should return true", func() {
				Expect(isOption).To(BeTrue())
			})
		})
		When("when token is semantics and is option flag", func() {
			var tokens = TokenList{}
			var tokenType TokenType = SemX2lktExplicitAssignment
			token := NewToken(0, tokenType, "-opt=value", tokens)
			tokens = append(tokens, token)
			isOptionPart := token.IsOptionFlag()
			It("should return true", func() {
				Expect(isOptionPart).To(BeTrue())
			})
		})
	})
	Describe("String func", func() {
		token := NewToken(0, CfTwoDashWord, "--two-dash", nil)
		It("should return a string", func() {
			Expect(token.String()).To(BeAssignableToTypeOf(""))
		})
	})
	Describe("MapToTypes func", func() {
		var tokens = TokenList{}
		tokens = append(tokens, NewToken(0, CfEndOfOptions, "--", tokens))
		tokens = append(tokens, NewToken(0, SemX2lktExplicitAssignment, "-opt=value", tokens))
		It("should return a list of token types", func() {
			Expect(tokens.MapToTypes()).To(BeAssignableToTypeOf([]TokenType{}))
			Expect(tokens.MapToTypes()).To(Equal([]TokenType{CfEndOfOptions, SemX2lktExplicitAssignment}))
		})
	})
})
