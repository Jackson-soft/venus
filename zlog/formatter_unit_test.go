package zlog

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TextFormatter", func() {
	Context("NewTextFmt", func() {
		It("should create formatter for valid level", func() {
			tf := NewTextFmt("info")
			Expect(tf).ShouldNot(BeNil())
			Expect(tf.tLevel).Should(Equal(InfoLevel))
		})

		It("should return nil for invalid level", func() {
			tf := NewTextFmt("invalid")
			Expect(tf).Should(BeNil())
		})
	})

	Context("SetLevel", func() {
		It("should change level successfully", func() {
			tf := NewTextFmt("info")
			err := tf.SetLevel("debug")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(tf.tLevel).Should(Equal(DebugLevel))
		})

		It("should return error for invalid level", func() {
			tf := NewTextFmt("info")
			err := tf.SetLevel("invalid")
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Format", func() {
		It("should return nil for level below threshold", func() {
			tf := NewTextFmt("warn")
			buf := tf.Format(InfoLevel, "should be filtered")
			Expect(buf).Should(BeNil())
		})

		It("should return nil for empty message", func() {
			tf := NewTextFmt("trace")
			buf := tf.Format(ErrorLevel, "")
			Expect(buf).Should(BeNil())
		})

		It("should format output with level and message", func() {
			tf := NewTextFmt("trace")
			buf := tf.Format(ErrorLevel, "test message")
			Expect(buf).ShouldNot(BeNil())
			output := string(buf)
			Expect(output).Should(ContainSubstring("error"))
			Expect(output).Should(ContainSubstring("test message"))
			Expect(output).Should(HaveSuffix("\n"))
		})
	})

	Context("WithFields", func() {
		It("should ignore empty fields", func() {
			tf := NewTextFmt("trace")
			tf.WithFields(Fields{})
			Expect(tf.data).Should(BeEmpty())
		})

		It("should add fields to output", func() {
			tf := NewTextFmt("trace")
			tf.WithFields(Fields{"key1": "val1", "key2": 42})
			Expect(tf.data).Should(HaveLen(2))
			buf := tf.Format(WarnLevel, "with fields")
			output := string(buf)
			Expect(output).Should(ContainSubstring("key1:val1"))
			Expect(output).Should(ContainSubstring("key2:42"))
		})

		It("should override existing fields", func() {
			tf := NewTextFmt("trace")
			tf.WithFields(Fields{"k": "old"})
			tf.WithFields(Fields{"k": "new"})
			Expect(tf.data["k"]).Should(Equal("new"))
		})
	})
})
