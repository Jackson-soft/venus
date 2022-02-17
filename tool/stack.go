package tool

import (
	"bytes"
	"runtime"
)

const (
	traceSize = 4096
	bitSize   = 10
)

// PanicTrace trace panic stack info.
func PanicTrace(kb int) []byte {
	if kb == 0 {
		kb = traceSize
	}

	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<bitSize) // 4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}

	return bytes.TrimRight(stack, "\n")
}
