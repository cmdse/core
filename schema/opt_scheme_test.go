package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OptionScheme", func() {
	Describe("SupportsTokenType method", func() {
		It("should return true when token type variant match one of OptionScheme", func() {
			Expect(OptionSchemePOSIXStrict.SupportsTokenType(SemPOSIXShortAssignmentLeftSide)).To(BeTrue())
		})
		It("should return false when token type variant does not match one of OptionScheme", func() {
			Expect(OptionSchemePOSIXStrict.SupportsTokenType(SemGNUSwitch)).To(BeFalse())
		})
	})
})
