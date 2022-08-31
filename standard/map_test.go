package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {
	Context("map test", func() {
		It("insert", func() {
			mm := standard.NewMap[int64, int64]()
			Expect(mm.Size()).Should(Equal(0))

			mm.Insert(43, 9)
			Expect(mm.Size()).Should(Equal(1))
		})

		It("erase", func() {
			mm := standard.NewMap[int64, string]()
			Expect(mm.Size()).Should(Equal(0))

			mm.Insert(43, "999")
			Expect(mm.Size()).Should(Equal(1))

			mm.Erase(43)
			Expect(mm.Size()).Should(Equal(0))
		})
	})
})
