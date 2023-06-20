package standard_test

import (
	"testing"

	"github.com/Jackson-soft/venus/standard"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("test queue", func() {
	queue := standard.NewQueue[int]()
	Context("test push", func() {
		It("queue", func() {
			queue.Push(1)
			Expect(queue.Size()).Should(Equal(uint(1)))

			queue.Push(2)
			Expect(queue.Size()).Should(Equal(uint(2)))

			queue.Push(3)
			Expect(queue.Size()).Should(Equal(uint(3)))
		})
	})

	Context("test pop", func() {
		It("pop", func() {
			v, ok := queue.Pop()
			Expect(queue.Size()).Should(Equal(uint(2)))
			Expect(v).Should(Equal(1))
			Expect(ok).Should(BeTrue())

			v, ok = queue.Pop()
			Expect(queue.Size()).Should(Equal(uint(1)))
			Expect(v).Should(Equal(2))
			Expect(ok).Should(BeTrue())

			v, ok = queue.Pop()
			Expect(queue.Size()).Should(Equal(uint(0)))
			Expect(v).Should(Equal(3))
			Expect(ok).Should(BeTrue())
		})
	})
})

func BenchmarkQueue2(b *testing.B) {
	q := standard.NewQueue[int]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push(3)
		}
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Pop()
		}
	})
}

func BenchmarkQueue(b *testing.B) {
	q := standard.NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}
