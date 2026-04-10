package zlog

import (
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	dir := b.TempDir()
	bb, err := NewInciseFile(dir, "", "", 0)
	if err != nil {
		b.Error(err)
	}
	msg := []byte("this is a message!!!\n")
	for range b.N {
		bb.Write(msg)
	}
}
