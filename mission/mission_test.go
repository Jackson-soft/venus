package mission_test

import (
	"fmt"
	"reflect"

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
