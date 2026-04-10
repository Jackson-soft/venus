package mission_test

import (
	"sync/atomic"
	"time"

	"github.com/Jackson-soft/venus/mission"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Producer", func() {
	Context("error cases", func() {
		It("should error for non-function handler", func() {
			err := mission.Instance().Producer("not a function")
			Expect(err).Should(Equal(mission.ErrTaskNotFunc))
		})

		It("should error for wrong param count (too few)", func() {
			fn := func(a int, b string) {}
			err := mission.Instance().Producer(fn, 1)
			Expect(err).Should(Equal(mission.ErrNumberOfParameters))
		})

		It("should error for wrong param count (too many)", func() {
			fn := func(a int, b string) {}
			err := mission.Instance().Producer(fn, 1, "hello", "extra")
			Expect(err).Should(Equal(mission.ErrNumberOfParameters))
		})
	})

	Context("no params", func() {
		It("should execute function with no params", func() {
			var called atomic.Int32
			fn := func() { called.Add(1) }

			err := mission.Instance().Producer(fn)
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(100 * time.Millisecond)
			Expect(called.Load()).Should(Equal(int32(1)))
		})
	})

	Context("with params", func() {
		It("should pass string param", func() {
			var result atomic.Value
			fn := func(name string) { result.Store(name) }

			err := mission.Instance().Producer(fn, "hello")
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(100 * time.Millisecond)
			Expect(result.Load()).Should(Equal("hello"))
		})

		It("should pass pointer param", func() {
			type Data struct{ Value int }
			var result atomic.Value
			fn := func(d *Data) { result.Store(d.Value) }

			err := mission.Instance().Producer(fn, &Data{Value: 42})
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(100 * time.Millisecond)
			Expect(result.Load()).Should(Equal(42))
		})
	})

	Context("multiple tasks", func() {
		It("should execute all tasks", func() {
			var count atomic.Int32
			fn := func(n int) { count.Add(1) }

			for i := range 10 {
				err := mission.Instance().Producer(fn, i)
				Expect(err).ShouldNot(HaveOccurred())
			}
			time.Sleep(200 * time.Millisecond)
			Expect(count.Load()).Should(Equal(int32(10)))
		})
	})

	Context("panic recovery", func() {
		It("should recover from handler panic and continue working", func() {
			fn := func() { panic("test panic") }
			err := mission.Instance().Producer(fn)
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(100 * time.Millisecond)

			var called atomic.Int32
			fn2 := func() { called.Add(1) }
			err = mission.Instance().Producer(fn2)
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(100 * time.Millisecond)
			Expect(called.Load()).Should(Equal(int32(1)))
		})
	})
})

var _ = Describe("Instance", func() {
	It("should return the same singleton", func() {
		Expect(mission.Instance()).Should(BeIdenticalTo(mission.Instance()))
	})
})
