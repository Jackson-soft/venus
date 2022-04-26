package standard_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Jackson-soft/venus/standard"
)

var _ = Describe("RingBuffer", func() {
	Context("one", func() {
		It("ring buffer", func() {
			buf := standard.NewRingBuf[int](5)
			data := 3
			buf.Push(data)

			Expect(buf.Pop()).Should(Equal(data))
		})
	})

	Context("full", func() {
		It("some", func() {
			buf := standard.NewRingBuf[int](3)
			for i := 0; i < 3; i++ {
				buf.Push(i)
			}

			Expect(buf.Pop()).Should(Equal(0))

			buf.Push(4)
			tt := buf.Pop()
			tt = buf.Pop()
			tt = buf.Pop()
			GinkgoWriter.Println(tt)
			buf.Push(6)
			tt = buf.Pop()
			ss := buf.Size()
			GinkgoWriter.Println(ss)
		})
	})
})
