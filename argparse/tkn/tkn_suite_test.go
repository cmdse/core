package tkn_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTkn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tkn Suite")
}
