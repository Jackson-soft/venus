package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jackson-soft/venus/standard"
)

func main() {
	mm := standard.NewMap[int, int]()
	mm.Insert(1, 11)
	mm.Insert(2, 22)

	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			go func(index int) {
				v, ok := mm.Find(1)
				log.Println(index, v, ok)
				ch <- 1
			}(i)
		}
	}()

	var num int
	for {
		select {
		case n := <-ch:
			num += n
			if num == 99 {
				fmt.Println(num)
				os.Exit(0)
			}
		}
	}
}
