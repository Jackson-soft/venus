package standard

import (
	"sort"
	"sync"
)

// 不重复切片
type Set struct {
	sync.RWMutex
	m_ map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		m_: make(map[string]struct{}),
	}
}

func (s *Set) Insert(key string) {
	s.Lock()
	s.m_[key] = struct{}{}
	s.Unlock()
}

func (s *Set) Erase(key string) {
	s.Lock()
	delete(s.m_, key)
	s.Unlock()
}

func (s *Set) Has(key string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m_[key]
	return ok
}

func (s *Set) Size() int {
	return len(s.m_)
}

func (s *Set) Empty() bool {
	return s.Size() == 0
}

func (s *Set) Clear() {
	s.Lock()
	s.m_ = make(map[string]struct{})
	s.Unlock()
}

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, s.Size())
	for i := range s.m_ {
		list = append(list, i)
	}
	return list
}

func (s *Set) SortList() []string {
	list := s.List()
	sort.Strings(list)
	return list
}
