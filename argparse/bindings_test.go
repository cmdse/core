package argparse

import (
	. "cmdse-cli/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bindings", func() {
	Describe("Contains func", func() {
		bindings := Bindings{BindLeft, BindRight}
		It("should return true when provided binding was given at initialization", func() {
			Expect(bindings.Contains(BindLeft)).To(BeTrue())
			Expect(bindings.Contains(BindRight)).To(BeTrue())
		})
		It("should return false when provided binding was not given at initialization", func() {
			Expect(bindings.Contains(BindNone)).To(BeFalse())
		})
	})
})
