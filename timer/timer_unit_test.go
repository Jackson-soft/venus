package timer_test

import (
	"sync/atomic"
	"time"

	"github.com/Jackson-soft/venus/timer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timer", func() {
	Context("single shot", func() {
		It("should fire exactly once", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var called atomic.Int32
			_, err := tt.Register(timer.Single, 200*time.Millisecond, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(500 * time.Millisecond)
			Expect(called.Load()).Should(Equal(int32(1)))
		})
	})

	Context("repetition", func() {
		It("should fire multiple times", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var called atomic.Int32
			_, err := tt.Register(timer.Repetition, 200*time.Millisecond, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(700 * time.Millisecond)
			Expect(called.Load()).Should(BeNumerically(">=", int32(2)))
		})
	})

	Context("Remove", func() {
		It("should prevent timer from firing", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var called atomic.Int32
			id, err := tt.Register(timer.Repetition, 200*time.Millisecond, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(tt.Remove(id)).Should(Succeed())
			time.Sleep(500 * time.Millisecond)
			Expect(called.Load()).Should(Equal(int32(0)))
		})

		It("should error for empty ID", func() {
			tt := timer.NewTimer(time.Second, 5)
			defer tt.Stop()
			Expect(tt.Remove("")).Should(HaveOccurred())
		})

		It("should error for non-existent ID", func() {
			tt := timer.NewTimer(time.Second, 5)
			defer tt.Stop()
			Expect(tt.Remove("non-existent-id")).Should(HaveOccurred())
		})
	})

	Context("Reset", func() {
		It("should restart the timer delay", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var called atomic.Int32
			id, err := tt.Register(timer.Single, 300*time.Millisecond, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())

			time.Sleep(100 * time.Millisecond)
			Expect(tt.Reset(id)).Should(Succeed())
			time.Sleep(250 * time.Millisecond)
			time.Sleep(200 * time.Millisecond)
			Expect(called.Load()).Should(Equal(int32(1)))
		})

		It("should error for empty ID", func() {
			tt := timer.NewTimer(time.Second, 5)
			defer tt.Stop()
			Expect(tt.Reset("")).Should(HaveOccurred())
		})

		It("should error for non-existent ID", func() {
			tt := timer.NewTimer(time.Second, 5)
			defer tt.Stop()
			Expect(tt.Reset("non-existent-id")).Should(HaveOccurred())
		})
	})

	Context("with args", func() {
		It("should pass args to handler", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var result atomic.Value
			_, err := tt.Register(timer.Single, 200*time.Millisecond, func(args any) {
				result.Store(args)
			}, "hello")
			Expect(err).ShouldNot(HaveOccurred())
			time.Sleep(500 * time.Millisecond)
			Expect(result.Load()).Should(Equal("hello"))
		})
	})

	Context("default values", func() {
		It("should handle zero tick and num", func() {
			tt := timer.NewTimer(0, 0)
			defer tt.Stop()

			var called atomic.Int32
			_, err := tt.Register(timer.Single, 2*time.Second, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Stop", func() {
		It("should stop firing after Stop", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)

			var called atomic.Int32
			_, err := tt.Register(timer.Repetition, 100*time.Millisecond, func(args any) {
				called.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())

			time.Sleep(300 * time.Millisecond)
			tt.Stop()
			countAfterStop := called.Load()
			time.Sleep(300 * time.Millisecond)
			Expect(called.Load()).Should(BeNumerically("<=", countAfterStop+1))
		})
	})

	Context("multiple registrations", func() {
		It("should handle multiple independent timers", func() {
			tt := timer.NewTimer(100*time.Millisecond, 10)
			defer tt.Stop()

			var count1, count2 atomic.Int32
			_, err := tt.Register(timer.Single, 200*time.Millisecond, func(args any) {
				count1.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())

			_, err = tt.Register(timer.Single, 300*time.Millisecond, func(args any) {
				count2.Add(1)
			}, nil)
			Expect(err).ShouldNot(HaveOccurred())

			time.Sleep(600 * time.Millisecond)
			Expect(count1.Load()).Should(Equal(int32(1)))
			Expect(count2.Load()).Should(Equal(int32(1)))
		})
	})
})
