package standard_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Jackson-soft/venus/standard"
)

var _ = Describe("Hashing", func() {
	It("hashing", func() {
		h := standard.NewHashing(3, nil)
		Expect(h).ShouldNot(BeNil())

		h.Add("11", "22", "33")

		str := h.Get("1") // 33
		Expect(str).ShouldNot(BeEmpty())
		str = h.Get("2") // 11
		Expect(str).ShouldNot(BeEmpty())
		str = h.Get("3") // 33
		Expect(str).ShouldNot(BeEmpty())

		h.Del("11")

		str = h.Get("1")
		Expect(str).ShouldNot(BeEmpty())
		str = h.Get("2")
		Expect(str).ShouldNot(BeEmpty())
		str = h.Get("3")
		Expect(str).ShouldNot(BeEmpty())
	})
})
