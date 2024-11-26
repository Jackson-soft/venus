package timer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Jackson-soft/venus/timer"
)

func TestRegister(t *testing.T) {
	tt := timer.NewTimer(time.Second, 5)
	tt.Register(timer.Single, 3*time.Second, func(args interface{}) {
		fmt.Println(args)
	}, "fsdfsdfsa")
}
