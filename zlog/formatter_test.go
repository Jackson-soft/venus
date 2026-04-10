package zlog

import "testing"

func BenchmarkFormattor(b *testing.B) {
	tf := TextFormatter{}
	for range b.N {
		tf.Format(InfoLevel, "this is a message!!!")
	}
}
