package standard

// 环形缓冲
type RingBuffer[T any] struct {
	size_    int
	write_   int // 读游标
	read_    int // 写游标
	element_ []T
}

func NewRingBuf[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		size_:    size,
		write_:   0,
		read_:    0,
		element_: make([]T, size),
	}
}

// 入队
func (r *RingBuffer[T]) Push(data T) bool {
	if (r.write_+1)%r.size_ == r.read_ {
		return false
	}
	r.element_[r.write_] = data
	r.write_ = (r.write_ + 1) % r.size_
	return true
}

// 出队
func (r *RingBuffer[T]) Pop() T {
	data := r.element_[r.read_]
	if r.read_ != r.write_ {
		r.read_ = (r.read_ + 1) % r.size_
	}
	return data
}

// 实际容纳的数据量
func (r *RingBuffer[T]) Size() int {
	return (r.write_ - r.read_ + r.size_) % r.size_
}

func (r *RingBuffer[T]) IsFull() bool {
	return r.read_ == r.write_
}
