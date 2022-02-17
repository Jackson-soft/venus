package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("test set", func() {
	s := standard.NewSet()

	It("test", func() {
		aa := s.List()

		Expect(len(aa)).Should(Equal(0))

		s.Insert("e")

		s.Insert("b")

		array := s.List()
		Expect(len(array)).Should(Equal(2))

		// sarray := s.SortList()

		s.Erase("b")

		s.Clear()
		Expect(s.Empty()).Should(BeTrue())
	})
})
