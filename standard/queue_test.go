package standard_test

import (
	"sync"
	"testing"

	"github.com/Jackson-soft/venus/standard"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := standard.NewQueue()
	q.Push(1)
	assert.Equal(t, q.Size(), uint(1))
	q.Push(2)
	assert.Equal(t, q.Size(), uint(2))
	q.Push(3)
	assert.Equal(t, q.Size(), uint(3))

	assert.Equal(t, q.Pop(), 1)
	assert.Equal(t, q.Size(), uint(2))
	assert.Equal(t, q.Pop(), 2)
	assert.Equal(t, q.Size(), uint(1))
	assert.Equal(t, q.Pop(), 3)
	assert.Equal(t, q.Size(), uint(0))
	assert.Equal(t, q.Pop(), nil)
	assert.Equal(t, q.Size(), uint(0))

	assert.Equal(t, q.Pop(), nil)

	q.Push(4)
	assert.Equal(t, q.Size(), uint(1))
	assert.Equal(t, q.Pop(), 4)
}

func TestQueue2(t *testing.T) {
	q := standard.NewQueue()

	var wg sync.WaitGroup
	wg.Add(2)

	total := 1000

	go func() {
		for i := 0; i < total; i++ {
			q.Push(i)
		}
	}()

	go func() {
		for i := 0; i < total; i++ {
			q.Pop()
		}
	}()

	wg.Done()

	t.Log(q.Size())
}

func BenchmarkQueue2(b *testing.B) {
	q := standard.NewQueue()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push(3)
		}
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Pop()
		}
	})
}

func BenchmarkQueue(b *testing.B) {
	q := standard.NewQueue()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}
