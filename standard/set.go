package standard

import (
	"sync"
)

// 不重复切片
type Set[T comparable] struct {
	mutex_ sync.RWMutex
	m_     map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		mutex_: sync.RWMutex{},
		m_:     make(map[T]struct{}),
	}
}

func (s *Set[T]) Insert(key T) {
	s.mutex_.Lock()
	s.m_[key] = struct{}{}
	s.mutex_.Unlock()
}

func (s *Set[T]) Erase(key T) {
	s.mutex_.Lock()
	delete(s.m_, key)
	s.mutex_.Unlock()
}

func (s *Set[T]) Has(key T) bool {
	s.mutex_.RLock()
	defer s.mutex_.RUnlock()
	_, ok := s.m_[key]
	return ok
}

func (s *Set[T]) Size() int {
	s.mutex_.RLock()
	defer s.mutex_.RUnlock()
	return len(s.m_)
}

func (s *Set[T]) Empty() bool {
	s.mutex_.RLock()
	defer s.mutex_.RUnlock()
	return s.Size() == 0
}

func (s *Set[T]) Clear() {
	s.mutex_.Lock()
	s.m_ = make(map[T]struct{})
	s.mutex_.Unlock()
}

func (s *Set[T]) List() []T {
	s.mutex_.RLock()
	defer s.mutex_.RUnlock()
	list := make([]T, 0, s.Size())
	for i := range s.m_ {
		list = append(list, i)
	}
	return list
}

/*
func (s *Set[T]) SortList() []T {
	list := s.List()
	sort.Strings(list)
	return list
}
*/
