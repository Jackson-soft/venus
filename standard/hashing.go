package standard

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 一致性hash

type HashFunc func(data []byte) uint32

type Hashing struct {
	hash_     HashFunc       // 自定义哈希算法，默认是crc32.ChecksumIEEE
	replicas_ int            // 虚拟节点倍数
	keys_     []int          // 哈希环,Sorted
	hashMap_  map[int]string // 虚拟节点与真实节点的映射表，键是虚拟节点的哈希值，值是真实节点的名称
}

func NewHashing(replicas int, fn HashFunc) *Hashing {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}

	return &Hashing{
		hash_:     fn,
		replicas_: replicas,
		keys_:     make([]int, 0),
		hashMap_:  make(map[int]string),
	}
}

// Add adds some keys to the hash.
func (h *Hashing) Add(keys ...string) {
	for i := range keys {
		for n := 0; n < h.replicas_; n++ {
			hash := int(h.hash_([]byte(strconv.Itoa(n) + keys[i])))
			h.keys_ = append(h.keys_, hash)
			h.hashMap_[hash] = keys[i]
		}
	}
	sort.Ints(h.keys_)
}

func (h *Hashing) Del(keys ...string) {
	for i := range keys {
		for n := 0; n < h.replicas_; n++ {
			hash := int(h.hash_([]byte(strconv.Itoa(n) + keys[i])))
			for m := range h.keys_ {
				if h.keys_[m] == hash {
					h.keys_ = append(h.keys_[:m], h.keys_[m+1:]...)
					break
				}
			}
			delete(h.hashMap_, hash)
		}
	}
	sort.Ints(h.keys_)
}

// Get gets the closest item in the hash to the provided key.
func (h *Hashing) Get(key string) string {
	if len(h.keys_) == 0 {
		return ""
	}

	hash := int(h.hash_([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(h.keys_), func(i int) bool {
		return h.keys_[i] >= hash
	})

	return h.hashMap_[h.keys_[idx%len(h.keys_)]]
}
