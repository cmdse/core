package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token", func() {
	Describe("IsBoundToOneOf method", func() {
		var tokens TokenList
		var tokenType TokenType = CfOneDashWord
		token := NewToken(0, tokenType, "-option", tokens)
		When("token type is context-free", func() {
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
		When("token type is semantic and bound 'None'", func() {
			var tokens TokenList
			var tokenType TokenType = SemX2lktSwitch
			token := NewToken(0, tokenType, "-option", tokens)
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
		When("token is context-free and has no semantic candidates", func() {
			var tokens TokenList
			var tokenType TokenType = CfTwoDashWord
			token := NewToken(0, tokenType, "--option", tokens)
			token.SemanticCandidates = nil
			Context("and provided bindings are 'None' and 'Left'", func() {
				boundToNoneOrLeft := token.IsBoundToOneOf(Bindings{BindNone, BindLeft})
				It("should return false", func() {
					Expect(boundToNoneOrLeft).To(BeFalse())
				})
			})
		})

	})
	Describe("IsBoundTo method", func() {
		var tokens = TokenList{}
		var tokenType TokenType = CfEndOfOptions
		token := NewToken(0, tokenType, "--", tokens)
		When("token type is semantic and bound to 'none'", func() {
			Context("and provided binding is 'None'", func() {
				It("should return true", func() {
					boundToNone := token.IsBoundTo(BindNone)
					Expect(boundToNone).To(BeTrue())
				})
			})
			Context("and provided binding is 'Left'", func() {
				It("should return false", func() {
					boundToLeft := token.IsBoundTo(BindLeft)
					Expect(boundToLeft).To(BeFalse())
				})
			})
		})
		When("token type is context-free and its semantic candidates are bound to 'None'", func() {
			token := NewToken(0, CfEndOfOptions, "-test", nil)
			When("provided with 'None' binding", func() {
				It("should return true", func() {
					boundToNone := token.IsBoundTo(BindNone)
					Expect(boundToNone).To(BeTrue())
				})
			})
			When("provided with 'Left' binding", func() {
				It("should return false", func() {
					boundToLeft := token.IsBoundTo(BindLeft)
					Expect(boundToLeft).To(BeFalse())
				})
			})

		})
		When("token type is context free and token has no semantic candidates", func() {
			var tokens TokenList
			var tokenType TokenType = CfTwoDashWord
			token := NewToken(0, tokenType, "-option", tokens)
			token.SemanticCandidates = nil
			Context("and provided bindings are 'None' and 'Left'", func() {
				boundToNoneOrLeft := token.IsBoundTo(BindNone)
				It("should return false", func() {
					Expect(boundToNoneOrLeft).To(BeFalse())
				})
			})
		})
	})
	Describe("IsOptionPart method", func() {
		checkIsOptionPart := func(tokenType TokenType, arg string) {
			token := NewToken(0, tokenType, arg, nil)
			isOption := token.IsOptionPart()
			It("should return true", func() {
				Expect(isOption).To(BeTrue())
			})
		}
		When("token is context free and its semantic candidates are option parts", func() {
			checkIsOptionPart(CfOneDashWord, "-test")
		})
		When("token is semantic and option part", func() {
			checkIsOptionPart(SemX2lktExplicitAssignment, "-opt=Value")
		})
	})
	Describe("IsOptionFlag method", func() {
		When("token is context free and its semantic candidates are option flags", func() {
			token := NewToken(0, CfEndOfOptions, "--", nil)
			isOption := token.IsOptionFlag()
			It("should return true", func() {
				Expect(isOption).To(BeTrue())
			})
		})
		When("token is semantics and is option flag", func() {
			token := NewToken(0, SemX2lktExplicitAssignment, "-opt=Value", nil)
			isOptionPart := token.IsOptionFlag()
			It("should return true", func() {
				Expect(isOptionPart).To(BeTrue())
			})
		})
	})
	Describe("String method", func() {
		token := NewToken(0, CfTwoDashWord, "--two-dash", nil)
		It("should return a string", func() {
			Expect(token.String()).To(BeAssignableToTypeOf(""))
		})
	})
	Describe("MapToTypes method", func() {
		var tokens = TokenList{}
		tokens = append(tokens, NewToken(0, CfEndOfOptions, "--", tokens))
		tokens = append(tokens, NewToken(0, SemX2lktExplicitAssignment, "-opt=Value", tokens))
		It("should return a list of token types", func() {
			Expect(tokens.MapToTypes()).To(BeAssignableToTypeOf([]TokenType{}))
			Expect(tokens.MapToTypes()).To(ConsistOf(CfEndOfOptions, SemX2lktExplicitAssignment))
		})
	})
	Describe("ReduceCandidatesWithScheme method", func() {
		When("given a CfWord", func() {
			token := NewToken(0, CfWord, "foo", nil)
			token.ReduceCandidatesWithScheme(OptSchemeXToolkitStrict)
			It("should keep semantic candidates", func() {
				Expect(token.SemanticCandidates).To(ContainElement(SemOperand))
			})
			It("should keep only value candidates conforming to the scheme", func() {
				Expect(token.SemanticCandidates).To(ConsistOf(SemOperand, SemX2lktImplicitAssignmentValue))
			})
		})
		When("given a flag Context-Free token", func() {
			token := NewToken(0, CfOneDashWordAlphaNum, "-foo", nil)
			token.ReduceCandidatesWithScheme(OptSchemeXToolkitStrict)
			It("should keep only flag candidates conforming to the scheme", func() {
				Expect(token.SemanticCandidates).To(ConsistOf(SemX2lktSwitch, SemX2lktImplicitAssignmentLeftSide))
			})
		})
	})

})
