package zlog

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Level", func() {
	Context("ParseLevel", func() {
		It("should parse all valid levels", func() {
			tests := []struct {
				input string
				want  Level
			}{
				{"trace", TraceLevel},
				{"debug", DebugLevel},
				{"info", InfoLevel},
				{"warn", WarnLevel},
				{"error", ErrorLevel},
				{"fatal", FatalLevel},
			}
			for _, tt := range tests {
				got, err := ParseLevel(tt.input)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(got).Should(Equal(tt.want))
			}
		})

		It("should be case insensitive", func() {
			got, err := ParseLevel("INFO")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(got).Should(Equal(InfoLevel))

			got, err = ParseLevel("Debug")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(got).Should(Equal(DebugLevel))
		})

		It("should return error for invalid levels", func() {
			for _, input := range []string{"unknown", "inf", "warning", "", "NULL"} {
				_, err := ParseLevel(input)
				Expect(err).Should(HaveOccurred())
			}
		})
	})

	Context("String", func() {
		It("should convert levels to strings", func() {
			Expect(TraceLevel.String()).Should(Equal("trace"))
			Expect(DebugLevel.String()).Should(Equal("debug"))
			Expect(InfoLevel.String()).Should(Equal("info"))
			Expect(WarnLevel.String()).Should(Equal("warn"))
			Expect(ErrorLevel.String()).Should(Equal("error"))
			Expect(FatalLevel.String()).Should(Equal("fatal"))
		})

		It("should return empty string for NULLLevel", func() {
			Expect(NULLLevel.String()).Should(BeEmpty())
		})
	})
})
