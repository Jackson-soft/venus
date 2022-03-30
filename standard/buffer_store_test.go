package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("buffer store", func() {
	var (
		buf *standard.Pool
		err error
	)
	Context("byte", func() {
		It("success", func() {
			buf, err = standard.NewBufferStore(2, 7)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(buf).ShouldNot(BeNil())

			data := buf.Get(3)
			Expect(data).ShouldNot(BeNil())

			buf.Put(data)
		})

		It("fail", func() {
			buf, err = standard.NewBufferStore(8, 1)
			Expect(err).Should(HaveOccurred())
			Expect(buf).Should(BeNil())
		})
	})
})
