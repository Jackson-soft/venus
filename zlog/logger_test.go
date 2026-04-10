package zlog

import (
	"testing"
)

func BenchmarkLoops(t *testing.B) {
	dir := t.TempDir()
	z := NewZLog(InfoLevel)
	b, err := NewInciseFile(dir, "xlog.log", "xxlog", 500)
	if err != nil {
		t.Error(err)
	}
	z.SetBackend(b)
	for i := 0; i < t.N; i++ {
		z.output(WarnLevel, "sdfasdf\n ssfdsdfs \n asdfsd")
	}
}
