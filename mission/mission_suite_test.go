package mission_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMission(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mission Suite")
}
