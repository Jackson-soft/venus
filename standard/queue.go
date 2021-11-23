package standard

import (
	"sync"
)

// 多线程队列
type Queue struct {
	mutex_ sync.Mutex
	size_  uint  // 容量
	head_  *node // 队头
	tail_  *node // 队尾
}

type node struct {
	value interface{}
	next  *node
}

func NewQueue() *Queue {
	return &Queue{
		size_: 0,
		head_: nil,
		tail_: nil,
	}
}

// 插入
func (q *Queue) Push(v interface{}) {
	n := &node{
		value: v,
	}

	q.mutex_.Lock()
	if q.head_ == nil {
		q.head_ = n
		q.tail_ = n
	} else {
		q.tail_.next = n
		q.tail_ = n
	}
	q.size_++
	q.mutex_.Unlock()
}

// 弹出
func (q *Queue) Pop() interface{} {
	q.mutex_.Lock()

	n := q.head_
	if n == nil {
		q.mutex_.Unlock()
		return nil
	}
	newHead := n.next
	if newHead == nil {
		q.tail_ = newHead
	}
	v := n.value
	q.head_ = newHead
	q.size_--
	q.mutex_.Unlock()
	return v
}

func (q *Queue) Size() uint {
	q.mutex_.Lock()
	defer q.mutex_.Unlock()
	return q.size_
}
