// Copyright 2019 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufpool

import "sync"

// GetFunc selects an buffer([]byte) from the Pool, removes it from the
// Pool, and returns it to the caller.
type GetFunc func() []byte

// PutFunc adds buffer([]byte) to the pool.
type PutFunc func(b []byte)

// A ByteArrayPool is a set of temporary buffer([]byte) that may be individually saved and
// retrieved.
type ByteArrayPool struct {
	pool *sync.Pool

	size        int
	preallocate int

	onceInit sync.Once
}

func (p *ByteArrayPool) init() {
	p.pool = &sync.Pool{New: func() interface{} {
		return make([]byte, 0, p.size)
	}}

	for i := 0; i < p.preallocate; i++ {
		p.pool.Put(p.pool.Get())
	}
}

// Get selects an buffer([]byte) from the Pool, removes it from the
// Pool, and returns it to the caller.
func (p *ByteArrayPool) Get() []byte {
	p.onceInit.Do(p.init)

	return p.pool.Get().([]byte)
}

// Put adds buffer([]byte) to the pool.
func (p *ByteArrayPool) Put(b []byte) {
	p.onceInit.Do(p.init)

	p.pool.Put(b[:0])
}

// NewByteArrayPool creates and initializes a new buffer([]byte) pool.
//  size: initial buffer size
//  preAllocate: count of preallocate
func NewByteArrayPool(size int, preallocate int) *ByteArrayPool {
	return &ByteArrayPool{size: size, preallocate: preallocate}
}
