package standard

import (
	"sync"
)

// 多线程队列
type Queue struct {
	mutex_ sync.RWMutex
	size_  uint  // 容量
	head_  *node // 队头
	tail_  *node // 队尾
}

type node struct {
	value_ interface{}
	next_  *node
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

// 弹出
func (q *Queue) Pop() interface{} {
	q.mutex_.Lock()
	defer q.mutex_.Unlock()

	if q.size_ == 0 {
		return nil
	}

	n := q.head_
	q.head_ = n.next_

	q.size_--
	return n.value_
}

func (q *Queue) Size() uint {
	q.mutex_.RLock()
	defer q.mutex_.RUnlock()
	return q.size_
}
