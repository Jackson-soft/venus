package zlog

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ZLog", func() {
	Context("creation", func() {
		It("should create a new ZLog", func() {
			z, _ := newTestZLog(DebugLevel)
			Expect(z).ShouldNot(BeNil())
			z.Stop()
		})
	})

	Context("level filtering", func() {
		It("should filter messages below threshold", func() {
			z, mb := newTestZLog(WarnLevel)
			z.Debugf("should not appear")
			time.Sleep(50 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(0))
			z.Stop()
		})

		It("should pass messages at or above threshold", func() {
			z, mb := newTestZLog(TraceLevel)
			z.Warnf("warning %d", 42)
			time.Sleep(50 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(1))
			z.Stop()
		})
	})

	Context("multiple backends", func() {
		It("should write to all backends", func() {
			z, mb1 := newTestZLog(TraceLevel)
			mb2 := &mockBackend{}
			z.AddBackend(mb2)

			z.Infof("multi backend")
			time.Sleep(50 * time.Millisecond)
			Expect(mb1.messageCount()).Should(Equal(1))
			Expect(mb2.messageCount()).Should(Equal(1))
			z.Stop()
		})
	})

	Context("SetLevel", func() {
		It("should change level dynamically", func() {
			z, mb := newTestZLog(TraceLevel)
			err := z.SetLevel("error")
			Expect(err).ShouldNot(HaveOccurred())

			z.Warnf("should be filtered now")
			time.Sleep(50 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(0))
			z.Stop()
		})

		It("should return error for invalid level", func() {
			z, _ := newTestZLog(TraceLevel)
			err := z.SetLevel("invalid")
			Expect(err).Should(HaveOccurred())
			z.Stop()
		})
	})

	Context("Sync", func() {
		It("should sync all backends", func() {
			z, mb := newTestZLog(TraceLevel)
			z.Infof("sync test")
			z.Sync()
			time.Sleep(50 * time.Millisecond)

			mb.mu.Lock()
			synced := mb.synced
			mb.mu.Unlock()
			Expect(synced).Should(BeTrue())
			z.Stop()
		})
	})

	Context("Stop", func() {
		It("should close all backends", func() {
			z, mb := newTestZLog(TraceLevel)
			z.Infof("before stop")
			time.Sleep(50 * time.Millisecond)
			z.Stop()
			time.Sleep(50 * time.Millisecond)

			mb.mu.Lock()
			closed := mb.closed
			mb.mu.Unlock()
			Expect(closed).Should(BeTrue())
		})
	})

	Context("WithFields", func() {
		It("should add fields to log output", func() {
			z, mb := newTestZLog(TraceLevel)
			z.WithFields(Fields{"user": "alice"}).Infof("hello")
			time.Sleep(50 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(1))
			z.Stop()
		})
	})

	Context("all log methods", func() {
		It("should output for all level methods", func() {
			z, mb := newTestZLog(TraceLevel)

			z.Tracef("trace %s", "msg")
			z.Debugf("debug %s", "msg")
			z.Infof("info %s", "msg")
			z.Warnf("warn %s", "msg")
			z.Errorf("error %s", "msg")
			z.Debugln("debugln")
			z.Infoln("infoln")
			z.Warnln("warnln")
			z.Errorln("errorln")

			time.Sleep(100 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(9))
			z.Stop()
		})
	})

	Context("SetFormattor", func() {
		It("should apply new formatter", func() {
			z, mb := newTestZLog(TraceLevel)
			tf := NewTextFmt("error")
			z.SetFormattor(tf)

			z.Infof("should be filtered by new formatter")
			time.Sleep(50 * time.Millisecond)
			Expect(mb.messageCount()).Should(Equal(0))
			z.Stop()
		})
	})

	Context("exported functions", func() {
		It("GetInstance should return non-nil", func() {
			Expect(GetInstance()).ShouldNot(BeNil())
		})
	})

	Context("file backend integration", func() {
		It("should write to file backend", func() {
			dir := GinkgoT().TempDir()
			z := NewZLog(InfoLevel)
			b, err := NewInciseFile(dir, "xlog.log", "xxlog", 500)
			Expect(err).ShouldNot(HaveOccurred())
			z.SetBackend(b)
			z.Infoln("this is a message!!")
			z.Stop()
		})

		It("should output to file backend", func() {
			dir := GinkgoT().TempDir()
			z := NewZLog(InfoLevel)
			b, err := NewInciseFile(dir, "xlog.log", "xxlog", 500)
			Expect(err).ShouldNot(HaveOccurred())
			z.SetBackend(b)
			z.output(WarnLevel, "sdfasdf\n ssfdsdfs \n asdfsd")
			z.Stop()
		})
	})
})
