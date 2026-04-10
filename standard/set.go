package standard

import (
	"cmp"
	"maps"
	"slices"
)

// Set 不重复集合
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
	clear(s.value_)
}

func (s *Set[T]) List() []T {
	return slices.Collect(maps.Keys(s.value_))
}

func (s *Set[T]) SortList() []T {
	return slices.Sorted(maps.Keys(s.value_))
}
