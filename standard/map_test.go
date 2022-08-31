package standard_test

import (
	"fmt"

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

		It("find", func() {
			mm := standard.NewMap[int64, int64]()

			mm.Insert(43, 666)
			Expect(mm.Size()).Should(Equal(1))

			vv, ok := mm.Find(43)
			Expect(ok).Should(BeTrue())
			Expect(vv).Should(Equal(int64(666)))

			vv, ok = mm.Find(222)
			Expect(ok).ShouldNot(BeTrue())
			Expect(vv).Should(Equal(int64(0)))
		})

		It("range", func() {
			mm := standard.NewMap[int64, int64]()

			mm.Insert(43, 666)
			Expect(mm.Size()).Should(Equal(1))
			mm.Insert(2, 55)

			mm.Range(func(key, value int64) bool {
				fmt.Println(key, value)
				return true
			})
		})
	})
})
