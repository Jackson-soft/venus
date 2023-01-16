package mission_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Jackson-soft/venus/mission"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mission", func() {
	Context("reflect", func() {
		It("function", func() {
			fn := func() {
				fmt.Println("ddff")
			}

			Expect(reflect.TypeOf(fn).Kind()).Should(Equal(reflect.Func))
		})
	})
})

func myTest(arg int) {
	fmt.Println(arg)
}

func BenchmarkBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mission.Instance().Producer(myTest, i)
	}
}
