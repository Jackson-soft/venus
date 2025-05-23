package standard

import (
	"errors"
	"math/bits"
	"sync"
)

/*
   内存池
*/

var (
	ErrSize = errors.New("max size can't be less than min size")
)

type (
	sizedPool struct {
		size int
		pool sync.Pool
	}

	// Pool is actually multiple pools which store buffers of specific size.
	// i.e. it can be three pools which return buffers 32K, 64K and 128K.
	Pool struct {
		minSize int
		maxSize int
		pools   []*sizedPool
	}
)

func newSizedPool(size int) *sizedPool {
	return &sizedPool{
		size: size,
		pool: sync.Pool{
			New: func() any { return makeSlicePointer(size) },
		},
	}
}

// New returns Pool which has buckets from minSize to maxSize.
// Buckets increase with the power of two, i.e with multiplier 2: [2b, 4b, 16b, ... , 1024b]
// Last pool will always be capped to maxSize.
func NewBufferStore(minSize, maxSize int) (*Pool, error) {
	if maxSize < minSize {
		return nil, ErrSize
	}
	const multiplier = 2
	pools := make([]*sizedPool, 0)
	curSize := minSize
	for curSize < maxSize {
		pools = append(pools, newSizedPool(curSize))
		curSize *= multiplier
	}
	pools = append(pools, newSizedPool(maxSize))
	return &Pool{
		minSize: minSize,
		maxSize: maxSize,
		pools:   pools,
	}, nil
}

func (p *Pool) findPool(size int) *sizedPool {
	if size > p.maxSize {
		return nil
	}
	div, rem := bits.Div64(0, uint64(size), uint64(p.minSize))
	idx := bits.Len64(div)
	if rem == 0 && div != 0 && (div&(div-1)) == 0 {
		idx--
	}
	return p.pools[idx]
}

// Get returns pointer to []byte which has len size.
// If there is no bucket with buffers >= size, slice will be allocated.
func (p *Pool) Get(size int) *[]byte {
	sp := p.findPool(size)
	if sp == nil {
		return makeSlicePointer(size)
	}
	buf, ok := sp.pool.Get().(*[]byte)
	if !ok {
		return makeSlicePointer(size)
	}
	*buf = (*buf)[:size]
	return buf
}

// Put returns pointer to slice to some bucket. Discards slice for which there is no bucket
func (p *Pool) Put(b *[]byte) {
	sp := p.findPool(cap(*b))
	if sp == nil {
		return
	}
	*b = (*b)[:cap(*b)]
	sp.pool.Put(b)
}

func makeSlicePointer(size int) *[]byte {
	data := make([]byte, size)
	return &data
}
