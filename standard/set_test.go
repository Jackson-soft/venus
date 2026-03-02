package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("test set", func() {
	ginkgo.Context("string set", func() {
		s := standard.NewSet[string]()

		ginkgo.It("test", func() {
			aa := s.List()

			gomega.Expect(aa).Should(gomega.BeEmpty())

			s.Insert("e")

			s.Insert("b")

			array := s.List()
			gomega.Expect(array).Should(gomega.HaveLen(2))

			// sarray := s.SortList()

			s.Erase("b")

			s.Clear()
			gomega.Expect(s.Empty()).Should(gomega.BeTrue())
		})
	})

	ginkgo.Context("int set", func() {
		s := standard.NewSet[int]()

		ginkgo.It("test", func() {
			aa := s.List()

			gomega.Expect(aa).Should(gomega.BeEmpty())

			s.Insert(1)

			s.Insert(4)

			s.Insert(2)

			array := s.List()
			gomega.Expect(array).Should(gomega.HaveLen(3))

			sarray := s.SortList()
			gomega.Expect(sarray).Should(gomega.Equal([]int{1, 2, 4}))

			s.Erase(4)

			s.Clear()
			gomega.Expect(s.Empty()).Should(gomega.BeTrue())
		})
	})
})
