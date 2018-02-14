package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProgramInterfaceModel", func() {
	Describe("Scheme method", func() {
		When("pim is null", func() {
			It("should not panic", func() {
				var pim *ProgramInterfaceModel
				Expect(func() { pim.Scheme() }).ToNot(Panic())
			})
			It("should return nil", func() {
				var pim *ProgramInterfaceModel
				Expect(pim.Scheme()).To(BeNil())
			})
		})
		When("pim is not null", func() {
			It("should return scheme", func() {
				var pim = NewProgramInterfaceModel(OptSchemeXToolkitStandard, nil)
				Expect(pim.Scheme()).To(Equal(OptSchemeXToolkitStandard))
			})
		})
	})
	Describe("DescriptionModel method", func() {
		When("pim is null", func() {
			It("should not panic", func() {
				var pim *ProgramInterfaceModel
				Expect(func() { pim.DescriptionModel() }).ToNot(Panic())
			})
			It("should return nil", func() {
				var pim *ProgramInterfaceModel
				Expect(pim.DescriptionModel()).To(BeNil())
			})
		})
		When("pim is not null", func() {
			It("should return option descriptions", func() {
				optDescriptions := OptDescriptionModel{}
				var pim = NewProgramInterfaceModel(OptSchemeXToolkitStandard, optDescriptions)
				Expect(pim.DescriptionModel()).To(Equal(optDescriptions))
			})
		})
	})
})
