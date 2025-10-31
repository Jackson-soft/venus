package tool_test

import (
	"github.com/Jackson-soft/venus/tool"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("aes test", func() {
	plaintext := []byte("this is massage")

	key := []byte("vYx3EXjCaSRe4QqWLn7Mpmcor0i2DdPw")
	ginkgo.Context("aes", func() {
		ginkgo.It("some", func() {
			data, err := tool.AesEncrypt(plaintext, key)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(data).ShouldNot(gomega.BeNil())

			txt, err := tool.AesDecrypt(data, key)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(txt).Should(gomega.Equal(string(plaintext)))
		})

		ginkgo.It("nil", func() {
			data, err := tool.AesEncrypt([]byte(""), key)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(data).Should(gomega.BeEmpty())

			txt, err := tool.AesDecrypt("", key)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Expect(txt).Should(gomega.Equal(""))
		})
	})
})
