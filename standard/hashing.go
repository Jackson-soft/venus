package standard

import (
	"hash/crc32"
	"slices"
	"strconv"
	"sync"
)

// 一致性hash

type (
	HashFunc func(data []byte) uint32

	Hashing struct {
		mu        sync.RWMutex
		hash_     HashFunc          // 自定义哈希算法，默认是crc32.ChecksumIEEE
		replicas_ int               // 虚拟节点倍数
		keys_     []uint32          // 哈希环,Sorted
		hashMap_  map[uint32]string // 虚拟节点与真实节点的映射表，键是虚拟节点的哈希值，值是真实节点的名称
	}
)

func NewHashing(replicas int, fn HashFunc) *Hashing {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}

	return &Hashing{
		hash_:     fn,
		replicas_: replicas,
		hashMap_:  make(map[uint32]string),
	}
}

// Add adds some keys to the hash.
func (h *Hashing) Add(keys ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, key := range keys {
		for n := range h.replicas_ {
			hash := h.hash_([]byte(strconv.Itoa(n) + key))
			h.keys_ = append(h.keys_, hash)
			h.hashMap_[hash] = key
		}
	}

	slices.Sort(h.keys_)
}

// Del removes keys from the hash.
func (h *Hashing) Del(keys ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	remove := make(map[uint32]struct{}, len(keys)*h.replicas_)

	for _, key := range keys {
		for n := range h.replicas_ {
			hash := h.hash_([]byte(strconv.Itoa(n) + key))
			remove[hash] = struct{}{}

			delete(h.hashMap_, hash)
		}
	}

	h.keys_ = slices.DeleteFunc(h.keys_, func(k uint32) bool {
		_, ok := remove[k]
		return ok
	})
}

// Get gets the closest item in the hash to the provided key.
func (h *Hashing) Get(key string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.keys_) == 0 {
		return ""
	}

	hash := h.hash_([]byte(key))
	idx, _ := slices.BinarySearch(h.keys_, hash)

	return h.hashMap_[h.keys_[idx%len(h.keys_)]]
}
