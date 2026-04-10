package standard_test

import (
	"github.com/Jackson-soft/venus/standard"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RingBuffer extended", func() {
	It("should start empty and not full", func() {
		buf := standard.NewRingBuf[int](4)
		Expect(buf.IsEmpty()).Should(BeTrue())
		Expect(buf.IsFull()).Should(BeFalse())
		Expect(buf.Size()).Should(Equal(0))
	})

	It("should detect full buffer", func() {
		buf := standard.NewRingBuf[int](4) // capacity = 3
		Expect(buf.Push(10)).Should(BeTrue())
		Expect(buf.Push(20)).Should(BeTrue())
		Expect(buf.Push(30)).Should(BeTrue())

		Expect(buf.Size()).Should(Equal(3))
		Expect(buf.IsFull()).Should(BeTrue())
		Expect(buf.IsEmpty()).Should(BeFalse())
		Expect(buf.Push(40)).Should(BeFalse()) // full
	})

	It("should pop in FIFO order", func() {
		buf := standard.NewRingBuf[string](4)
		buf.Push("a")
		buf.Push("b")
		buf.Push("c")
		Expect(buf.Pop()).Should(Equal("a"))
		Expect(buf.Pop()).Should(Equal("b"))
		Expect(buf.Pop()).Should(Equal("c"))
		Expect(buf.IsEmpty()).Should(BeTrue())
	})

	It("should handle wrap-around correctly", func() {
		buf := standard.NewRingBuf[int](3) // capacity = 2
		buf.Push(1)
		buf.Push(2)
		buf.Pop()    // removes 1
		buf.Push(3)  // wraps around
		Expect(buf.Pop()).Should(Equal(2))
		Expect(buf.Pop()).Should(Equal(3))
		Expect(buf.IsEmpty()).Should(BeTrue())
	})
})

var _ = Describe("Map extended", func() {
	It("should clear all entries", func() {
		mm := standard.NewMap[string, int]()
		mm.Insert("a", 1)
		mm.Insert("b", 2)
		Expect(mm.Size()).Should(Equal(2))

		mm.Clear()
		Expect(mm.Size()).Should(Equal(0))
		_, ok := mm.Find("a")
		Expect(ok).Should(BeFalse())
	})
})

var _ = Describe("Set extended", func() {
	It("should check existence", func() {
		s := standard.NewSet[string]()
		s.Insert("hello")
		s.Insert("world")
		Expect(s.Exist("hello")).Should(BeTrue())
		Expect(s.Exist("notfound")).Should(BeFalse())
	})

	It("should ignore duplicate inserts for size", func() {
		s := standard.NewSet[int]()
		s.Insert(1)
		s.Insert(2)
		s.Insert(1)
		Expect(s.Size()).Should(Equal(2))
	})
})

var _ = Describe("Queue extended", func() {
	It("should return false for pop on empty queue", func() {
		q := standard.NewQueue[int]()
		v, ok := q.Pop()
		Expect(ok).Should(BeFalse())
		Expect(v).Should(Equal(0))
	})

	It("should return false for front on empty queue", func() {
		q := standard.NewQueue[string]()
		v, ok := q.Front()
		Expect(ok).Should(BeFalse())
		Expect(v).Should(BeEmpty())
	})

	It("should push and pop in FIFO order", func() {
		q := standard.NewQueue[int]()
		q.Push(10)
		q.Push(20)
		q.Push(30)
		Expect(q.Size()).Should(Equal(uint(3)))

		v, ok := q.Front()
		Expect(ok).Should(BeTrue())
		Expect(v).Should(Equal(10))

		v, ok = q.Pop()
		Expect(ok).Should(BeTrue())
		Expect(v).Should(Equal(10))
		v, _ = q.Pop()
		Expect(v).Should(Equal(20))
		v, _ = q.Pop()
		Expect(v).Should(Equal(30))
	})

	It("should erase head safely including on empty", func() {
		q := standard.NewQueue[int]()
		q.Push(1)
		q.Push(2)
		q.Erase()
		Expect(q.Size()).Should(Equal(uint(1)))
		q.Erase()
		q.Erase() // extra erase on empty is safe
		Expect(q.Size()).Should(Equal(uint(0)))
	})
})

var _ = Describe("Hashing extended", func() {
	It("should return empty string for empty hashing", func() {
		h := standard.NewHashing(3, nil)
		Expect(h.Get("anything")).Should(BeEmpty())
	})

	It("should be consistent for the same key", func() {
		h := standard.NewHashing(50, nil)
		h.Add("server1", "server2", "server3")
		first := h.Get("mykey")
		for range 100 {
			Expect(h.Get("mykey")).Should(Equal(first))
		}
	})

	It("should still resolve after deleting a node", func() {
		h := standard.NewHashing(10, nil)
		h.Add("a", "b", "c")
		h.Del("a")
		got := h.Get("testkey")
		Expect(got).ShouldNot(BeEmpty())
	})
})

var _ = Describe("BufferStore extended", func() {
	It("should get and put buffers", func() {
		p, err := standard.NewBufferStore(4, 64)
		Expect(err).ShouldNot(HaveOccurred())

		buf := p.Get(10)
		Expect(buf).ShouldNot(BeNil())
		Expect(*buf).Should(HaveLen(10))
		p.Put(buf)
	})

	It("should handle oversize requests", func() {
		p, err := standard.NewBufferStore(4, 16)
		Expect(err).ShouldNot(HaveOccurred())

		buf := p.Get(32)
		Expect(buf).ShouldNot(BeNil())
		Expect(*buf).Should(HaveLen(32))
		p.Put(buf)
	})

	It("should error when maxSize < minSize", func() {
		_, err := standard.NewBufferStore(64, 4)
		Expect(err).Should(HaveOccurred())
	})
})
