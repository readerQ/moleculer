package transit_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTransit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transit Suite")
}
