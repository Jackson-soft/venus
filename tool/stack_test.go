package tool_test

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/Jackson-soft/venus/tool"
)

func TestStack(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("&&&&&")
			fmt.Println(err)
			fmt.Println("=========")
			buf := debug.Stack()
			fmt.Println(string(buf))

			fmt.Println("----------")
			buf = tool.PanicTrace(0)
			fmt.Println(string(buf))

			fmt.Println("******")
		}
	}()

	tt := []int{1, 3, 4}
	t.Log(tt[9])
}
