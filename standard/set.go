package standard

import (
	"cmp"
	"slices"
)

// 不重复切片
type Set[T cmp.Ordered] struct {
	value_ map[T]struct{}
}

func NewSet[T cmp.Ordered]() Set[T] {
	return Set[T]{
		value_: make(map[T]struct{}),
	}
}

func (s *Set[T]) Insert(key T) {
	s.value_[key] = struct{}{}
}

func (s *Set[T]) Erase(key T) {
	delete(s.value_, key)
}

func (s *Set[T]) Exist(key T) bool {
	_, ok := s.value_[key]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.value_)
}

func (s *Set[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Clear() {
	s.value_ = make(map[T]struct{})
}

func (s *Set[T]) List() []T {
	list := make([]T, s.Size())
	i := 0
	for key := range s.value_ {
		list[i] = key
		i++
	}
	return list
}

func (s *Set[T]) SortList() []T {
	list := s.List()
	slices.Sort(list)
	return list
}
