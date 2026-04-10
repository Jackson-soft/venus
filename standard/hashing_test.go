package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hashing", func() {
	var h *standard.Hashing

	BeforeEach(func() {
		h = standard.NewHashing(3, nil)
	})

	It("should create a valid instance", func() {
		Expect(h).ShouldNot(BeNil())
	})

	Context("with no nodes", func() {
		It("should return empty string", func() {
			Expect(h.Get("anything")).Should(BeEmpty())
		})
	})

	Context("with nodes added", func() {
		BeforeEach(func() {
			h.Add("node1", "node2", "node3")
		})

		It("should return a node for any key", func() {
			Expect(h.Get("key1")).ShouldNot(BeEmpty())
			Expect(h.Get("key2")).ShouldNot(BeEmpty())
			Expect(h.Get("key3")).ShouldNot(BeEmpty())
		})

		It("should consistently map the same key to the same node", func() {
			first := h.Get("key1")
			Expect(h.Get("key1")).Should(Equal(first))
			Expect(h.Get("key1")).Should(Equal(first))
		})

		Context("after deleting a node", func() {
			BeforeEach(func() {
				h.Del("node1")
			})

			It("should still return results for all keys", func() {
				Expect(h.Get("key1")).ShouldNot(BeEmpty())
				Expect(h.Get("key2")).ShouldNot(BeEmpty())
				Expect(h.Get("key3")).ShouldNot(BeEmpty())
			})

			It("should not return the deleted node", func() {
				for _, key := range []string{"a", "b", "c", "d", "e"} {
					Expect(h.Get(key)).Should(BeElementOf("node2", "node3"))
				}
			})
		})
	})

	Context("with custom hash function", func() {
		It("should use the provided function", func() {
			custom := standard.NewHashing(1, func(data []byte) uint32 {
				return uint32(len(data))
			})
			custom.Add("a", "bb")
			Expect(custom.Get("test")).ShouldNot(BeEmpty())
		})
	})
})
