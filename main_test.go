package optioner_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFuncs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Optioner Suite")
}
