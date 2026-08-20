package tool_test

import (
	"github.com/Jackson-soft/venus/tool"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sha3", func() {
	It("should be deterministic", func() {
		Expect(tool.Sha3("hello")).Should(Equal(tool.Sha3("hello")))
	})

	It("should differ for different inputs", func() {
		Expect(tool.Sha3("hello")).ShouldNot(Equal(tool.Sha3("world")))
	})

	It("should match known SHA3-256 value", func() {
		expected := "3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392"
		Expect(tool.Sha3("hello")).Should(Equal(expected))
	})

	It("should produce non-empty hash for empty string", func() {
		Expect(tool.Sha3("")).ShouldNot(BeEmpty())
	})
})

var _ = Describe("PKCS7", func() {
	Context("Padding and UnPadding", func() {
		It("should round-trip correctly", func() {
			data := []byte("hello")
			padded := tool.PKCS7Padding(data, 16)
			Expect(len(padded) % 16).Should(Equal(0))

			unpadded, err := tool.PKCS7UnPadding(padded)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(unpadded)).Should(Equal("hello"))
		})
	})

	Context("UnPadding errors", func() {
		It("should error for empty input", func() {
			_, err := tool.PKCS7UnPadding([]byte{})
			Expect(err).Should(HaveOccurred())
		})

		It("should error for invalid padding byte", func() {
			_, err := tool.PKCS7UnPadding([]byte{0xFF})
			Expect(err).Should(HaveOccurred())
		})

		It("should error for zero padding byte", func() {
			_, err := tool.PKCS7UnPadding([]byte{0x00})
			Expect(err).Should(HaveOccurred())
		})
	})
})

var _ = Describe("AES", func() {
	key32 := []byte("0123456789abcdef0123456789abcdef")
	key16 := []byte("0123456789abcdef")

	Context("encrypt and decrypt", func() {
		It("should round-trip successfully", func() {
			plaintext := "this is a secret message"
			encrypted, err := tool.AesEncrypt([]byte(plaintext), key32)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(encrypted).ShouldNot(BeEmpty())
			Expect(encrypted).ShouldNot(Equal(plaintext))

			decrypted, err := tool.AesDecrypt(encrypted, key32)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(decrypted).Should(Equal(plaintext))
		})
	})

	Context("empty input", func() {
		It("should return empty for encrypt", func() {
			encrypted, err := tool.AesEncrypt([]byte(""), key16)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(encrypted).Should(BeEmpty())
		})

		It("should return empty for decrypt", func() {
			decrypted, err := tool.AesDecrypt("", key16)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(decrypted).Should(BeEmpty())
		})
	})

	Context("errors", func() {
		It("should error for invalid key size on encrypt", func() {
			_, err := tool.AesEncrypt([]byte("data"), []byte("short"))
			Expect(err).Should(HaveOccurred())
		})

		It("should error for invalid key size on decrypt", func() {
			_, err := tool.AesDecrypt("aabbccdd", []byte("short"))
			Expect(err).Should(HaveOccurred())
		})

		It("should error for invalid hex on decrypt", func() {
			_, err := tool.AesDecrypt("not-hex!", key16)
			Expect(err).Should(HaveOccurred())
		})
	})
})

var _ = Describe("Unique", func() {
	It("should remove duplicates preserving order", func() {
		Expect(tool.Unique([]int{1, 2, 2, 3, 1, 4})).Should(Equal([]int{1, 2, 3, 4}))
	})

	It("should handle empty slice", func() {
		Expect(tool.Unique([]int{})).Should(BeEmpty())
	})

	It("should handle all same elements", func() {
		Expect(tool.Unique([]int{5, 5, 5})).Should(Equal([]int{5}))
	})

	It("should handle no duplicates", func() {
		Expect(tool.Unique([]int{1, 2, 3})).Should(Equal([]int{1, 2, 3}))
	})

	It("should work with strings", func() {
		Expect(tool.Unique([]string{"a", "b", "a", "c", "b"})).Should(Equal([]string{"a", "b", "c"}))
	})
})

type person struct {
	ID   int
	Name string
}

var _ = Describe("UniqueFunc", func() {
	It("should remove duplicates using custom equal function by field", func() {
		items := []person{
			{ID: 1, Name: "alice"},
			{ID: 2, Name: "bob"},
			{ID: 1, Name: "alice2"},
		}
		result := tool.UniqueFunc(items, func(a, b person) bool { return a.ID == b.ID })
		Expect(result).Should(HaveLen(2))
		Expect(result[0].ID).Should(Equal(1))
		Expect(result[1].ID).Should(Equal(2))
	})

	It("should preserve order", func() {
		items := []person{
			{ID: 3, Name: "c"},
			{ID: 1, Name: "a"},
			{ID: 3, Name: "c2"},
			{ID: 2, Name: "b"},
		}
		result := tool.UniqueFunc(items, func(a, b person) bool { return a.ID == b.ID })
		Expect(result).Should(Equal([]person{
			{ID: 3, Name: "c"},
			{ID: 1, Name: "a"},
			{ID: 2, Name: "b"},
		}))
	})

	It("should handle empty slice", func() {
		result := tool.UniqueFunc([]person{}, func(a, b person) bool { return a.ID == b.ID })
		Expect(result).Should(BeEmpty())
	})

	It("should work with comparable types too", func() {
		result := tool.UniqueFunc([]int{1, 2, 2, 3}, func(a, b int) bool { return a == b })
		Expect(result).Should(Equal([]int{1, 2, 3}))
	})
})

var _ = Describe("UniqueBy", func() {
	It("should remove duplicates by key preserving order", func() {
		items := []person{
			{ID: 1, Name: "alice"},
			{ID: 2, Name: "bob"},
			{ID: 1, Name: "alice2"},
		}
		result := tool.UniqueBy(items, func(p person) int { return p.ID })
		Expect(result).Should(Equal([]person{
			{ID: 1, Name: "alice"},
			{ID: 2, Name: "bob"},
		}))
	})

	It("should work with string keys", func() {
		items := []person{
			{ID: 1, Name: "alice"},
			{ID: 2, Name: "alice"},
			{ID: 3, Name: "bob"},
		}
		result := tool.UniqueBy(items, func(p person) string { return p.Name })
		Expect(result).Should(HaveLen(2))
		Expect(result[0].Name).Should(Equal("alice"))
		Expect(result[1].Name).Should(Equal("bob"))
	})

	It("should handle empty slice", func() {
		result := tool.UniqueBy([]person{}, func(p person) int { return p.ID })
		Expect(result).Should(BeEmpty())
	})

	It("should work with comparable types", func() {
		result := tool.UniqueBy([]int{1, 2, 2, 3, 1}, func(v int) int { return v })
		Expect(result).Should(Equal([]int{1, 2, 3}))
	})
})

var _ = Describe("RemoveElements", func() {
	It("should return source when toRemove is empty", func() {
		Expect(tool.RemoveElements([]int{1, 2, 3}, []int{})).Should(Equal([]int{1, 2, 3}))
	})

	It("should remove specified elements", func() {
		Expect(tool.RemoveElements([]int{1, 2, 3, 4, 5}, []int{2, 4})).Should(Equal([]int{1, 3, 5}))
	})

	It("should remove all elements", func() {
		Expect(tool.RemoveElements([]int{1, 2, 3}, []int{1, 2, 3})).Should(BeEmpty())
	})

	It("should handle non-existent elements", func() {
		Expect(tool.RemoveElements([]int{1, 2}, []int{3, 4})).Should(Equal([]int{1, 2}))
	})

	It("should handle empty source", func() {
		Expect(tool.RemoveElements([]int{}, []int{1, 2})).Should(BeEmpty())
	})
})

var _ = Describe("PanicTrace", func() {
	It("should capture panic stack trace", func() {
		defer func() {
			if r := recover(); r != nil {
				buf := tool.PanicTrace(0)
				Expect(buf).ShouldNot(BeEmpty())
			}
		}()
		var s []int
		_ = s[100]
	})

	It("should capture with custom size", func() {
		defer func() {
			if r := recover(); r != nil {
				buf := tool.PanicTrace(8)
				Expect(buf).ShouldNot(BeEmpty())
			}
		}()
		panic("test panic")
	})
})
