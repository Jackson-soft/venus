package timer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timer Suite")
}
