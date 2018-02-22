package schema_test

import (
	"github.com/cmdse/core/argparse"
	"github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("[regression test, issue #1] TokenList#CheckEndOfOptions and token#ReduceCandidatesWithScheme result in unexpectedly large Token#SemanticCandidates slices", func() {
	It("should not reproduce bug", func() {
		args := []string{"-option", "-long-option", "--", "-arg", "--arg2", "argument"}
		tokens := argparse.InitTokens(args)
		problematicToken := tokens[3]
		tokens.CheckEndOfOptions()
		lengthAfterEndOfOpt := len(problematicToken.SemanticCandidates)
		problematicToken.ReduceCandidatesWithScheme(schema.OptSchemeXToolkitStrict)
		Expect(len(problematicToken.SemanticCandidates)).To(BeNumerically("<=", lengthAfterEndOfOpt))
	})
})
