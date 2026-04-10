package httpkit_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHttpkit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httpkit Suite")
}
