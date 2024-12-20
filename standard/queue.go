package standard

import (
	"sync"
)

// 多线程队列
type Queue[T any] struct {
	mutex_ sync.RWMutex
	size_  uint     // 容量
	head_  *node[T] // 队头
	tail_  *node[T] // 队尾
}

type node[T any] struct {
	value_ T
	next_  *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		size_: 0,
		head_: nil,
		tail_: nil,
	}
}

// 插入
func (q *Queue[T]) Push(v T) {
	n := &node[T]{
		value_: v,
		next_:  nil,
	}

	q.mutex_.Lock()
	if q.size_ == 0 {
		q.head_ = n
	} else {
		q.tail_.next_ = n
	}
	q.tail_ = n
	q.size_++
	q.mutex_.Unlock()
}

// 访问头部
func (q *Queue[T]) Front() (T, bool) {
	q.mutex_.RLock()
	defer q.mutex_.RUnlock()

	var value T
	if q.size_ == 0 {
		return value, false
	}
	return q.head_.value_, true
}

// 弹出头部，并删除
func (q *Queue[T]) Pop() (T, bool) {
	q.mutex_.Lock()
	defer q.mutex_.Unlock()

	var value T
	if q.size_ == 0 {
		return value, false
	}

	n := q.head_
	q.head_ = n.next_

	q.size_--
	return n.value_, true
}

// 删除头部
func (q *Queue[T]) Erase() {
	q.mutex_.Lock()
	defer q.mutex_.Unlock()

	if q.size_ > 0 {
		n := q.head_
		q.head_ = n.next_

		q.size_--
	}
}

func (q *Queue[T]) Size() uint {
	q.mutex_.RLock()
	defer q.mutex_.RUnlock()
	return q.size_
}
