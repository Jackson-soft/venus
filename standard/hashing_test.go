package standard_test

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/Jackson-soft/venus/standard"
)

var _ = ginkgo.Describe("Hashing", func() {
	ginkgo.It("hashing", func() {
		h := standard.NewHashing(3, nil)
		gomega.Expect(h).ShouldNot(gomega.BeNil())

		h.Add("11", "22", "33")

		str := h.Get("1") // 33
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())
		str = h.Get("2") // 11
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())
		str = h.Get("3") // 33
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())

		h.Del("11")

		str = h.Get("1")
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())
		str = h.Get("2")
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())
		str = h.Get("3")
		gomega.Expect(str).ShouldNot(gomega.BeEmpty())
	})
})
