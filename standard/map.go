package standard

import "sync"

type Map[K comparable, V any] struct {
	mutex_ sync.RWMutex
	value_ map[K]V
}

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		mutex_: sync.RWMutex{},
		value_: make(map[K]V),
	}
}

func (m *Map[K, V]) Size() int {
	m.mutex_.RLock()
	defer m.mutex_.RUnlock()

	return len(m.value_)
}

func (m *Map[K, V]) Clear() {
	m.mutex_.Lock()
	m.value_ = make(map[K]V)
	m.mutex_.Unlock()
}

func (m *Map[K, V]) Insert(key K, value V) {
	m.mutex_.Lock()
	m.value_[key] = value
	m.mutex_.Unlock()
}

func (m *Map[K, V]) Erase(key K) {
	m.mutex_.Lock()
	delete(m.value_, key)
	m.mutex_.Unlock()
}
