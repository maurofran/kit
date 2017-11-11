package assert_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAssert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assert Suite")
}
