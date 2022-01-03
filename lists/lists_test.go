package lists_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lists Suite")
}
