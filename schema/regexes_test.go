package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParametricRegexBuilder", func() {

	Describe("assembleString method", func() {
		It("should assemble provided flagName and paramName", func() {
			regextest := &ParametricRegexBuilder{
				`^--(%s)=(%s)$`,
				"A",
				"B",
			}
			answer, _ := regextest.assembleString("N", "P")
			Expect(answer).To(Equal(`^--(N)=(P)$`))
		})
	})
	Describe("BuildDefault method", func() {
		It("should assemble inner flagName and inner paramName into a regex", func() {
			regextest := &ParametricRegexBuilder{
				`^--(%s)=(%s)$`,
				"A",
				"B",
			}
			answer := regextest.BuildDefault()
			Expect(answer.MatchString(`--A=B`)).To(BeTrue())
		})
	})
	Describe("Build method", func() {
		It("should assemble provided flagName and paramName into a regex", func() {
			regextest := &ParametricRegexBuilder{
				`^--(%s)=(%s)$`,
				"A",
				"B",
			}
			answer, _ := regextest.Build("N", "P")
			Expect(answer.MatchString(`--N=P`)).To(BeTrue())
		})
		It("should assemble provided flagName with inner paramName into a regex when provided paramName is empty", func() {
			regextest := &ParametricRegexBuilder{
				`^--(%s)=(%s)$`,
				"A",
				"B",
			}
			answer, _ := regextest.Build("N", "")
			Expect(answer.MatchString(`--N=B`)).To(BeTrue())
		})
		It("should return an error when provided flagName is empty", func() {
			regextest := &ParametricRegexBuilder{
				`^--(%s)=(%s)$`,
				"A",
				"B",
			}
			_, err := regextest.Build("", "B")
			Expect(err).To(HaveOccurred())
		})
	})
})
