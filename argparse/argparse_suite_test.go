package argparse_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArgparse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Argparse Suite")
}
