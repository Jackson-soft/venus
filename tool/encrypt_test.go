package tool_test

import (
	"github.com/Jackson-soft/venus/tool"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("aes test", func() {
	plaintext := []byte("this is massage")

	key := []byte("vYx3EXjCaSRe4QqWLn7Mpmcor0i2DdPw")

	It("df", func() {
		data, err := tool.AesEncrypt(plaintext, key)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(data).ShouldNot(BeNil())

		txt, err := tool.AesDecrypt(data, key)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(txt).Should(Equal(string(plaintext)))
	})
})
