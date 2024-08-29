package standard_test

import (
	"fmt"

	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("test set", func() {
	Context("string set", func() {
		s := standard.NewSet[string]()

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

	Context("int set", func() {
		s := standard.NewSet[int]()

		It("test", func() {
			aa := s.List()

			Expect(len(aa)).Should(Equal(0))

			s.Insert(1)

			s.Insert(4)

			s.Insert(2)

			array := s.List()
			Expect(len(array)).Should(Equal(3))

			sarray := s.SortList()
			fmt.Printf("%v\n", sarray)

			s.Erase(4)

			s.Clear()
			Expect(s.Empty()).Should(BeTrue())
		})
	})
})
