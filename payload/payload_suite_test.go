package payload_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payload Suite")
}
