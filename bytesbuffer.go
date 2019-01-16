// Copyright 2019 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufpool

import (
	"bytes"
	"sync"
)

// A BytesBufferPool is a set of temporary bytes.Buffer
// that may be individually saved and retrieved.
type BytesBufferPool struct {
	pool *ByteArrayPool

	onceInit sync.Once
}

func (p *BytesBufferPool) init() {
	if p.pool == nil {
		p.pool = &ByteArrayPool{}
	}
}

// Get selects an bytes.Buffer from the Pool, removes it from the
// Pool, and returns it to the caller.
func (p *BytesBufferPool) Get() *bytes.Buffer {
	p.onceInit.Do(p.init)

	return bytes.NewBuffer(p.pool.Get())
}

// Put adds bytes.Buffer to the pool.
func (p *BytesBufferPool) Put(b *bytes.Buffer) {
	p.onceInit.Do(p.init)

	b.Reset()
	p.pool.Put(b.Bytes())
}

// NewBytesBufferPool creates and initializes a new bytes.Buffer pool.
//  size: initial buffer size
//  preAllocate: count of preallocate
func NewBytesBufferPool(size int, preallocate int) *BytesBufferPool {
	return &BytesBufferPool{pool: NewByteArrayPool(size, preallocate)}
}
