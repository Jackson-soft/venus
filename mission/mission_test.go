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

func BenchmarkBuf(b *testing.B) {
	fun := func(arg int) {
		_, _ = fmt.Println(arg)
	}
	for i := 0; i < b.N; i++ {
		mission.Instance().Producer(fun, i)
	}
}

type mm struct {
	Age  int64
	Name string
}

func pmm(arg *mm) {
	fmt.Println(arg.Age, arg.Name)
}

func pval(arg int64) {
	fmt.Println(arg)
}

func TestTask(t *testing.T) {
	for i := 0; i < 10; i++ {
		mission.Instance().Producer(pmm, &mm{
			Age:  24,
			Name: "4dd433",
		})

		mission.Instance().Producer(pval, int64(888))
	}
}

func TestValue(t *testing.T) {
	arg := &mm{
		Age:  10,
		Name: "test",
	}

	val := reflect.ValueOf(arg)

	t.Log(val)

	fnm := reflect.ValueOf(pmm)

	fnm.Call([]reflect.Value{val})

	fnv := reflect.ValueOf(pval)

	valv := reflect.ValueOf(int64(555))

	fnv.Call([]reflect.Value{valv})
}
