// Copyright 2019 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufpool

import (
	"testing"
)

func TestBytesBufferPool_Get(t *testing.T) {
	type args struct {
		size        int
		preallocate int
	}
	tests := []struct {
		name   string
		fields args
	}{
		{name: "size: 4096, preallocate: 0", fields: args{size: 4096, preallocate: 0}},
		{name: "size: 4096, preallocate: 4", fields: args{size: 4096, preallocate: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BytesBufferPool{
				pool: NewByteArrayPool(tt.fields.size, tt.fields.preallocate),
			}
			got := p.Get()
			if c := cap(got.Bytes()); c != tt.fields.size {
				t.Errorf("cap(bufpool.Get()) = %d, want %d", c, tt.fields.size)
			}
			if l := len(got.Bytes()); l != 0 {
				t.Errorf("len(bufpool.Get()) = %d, want 0", l)
			}
		})
	}
}

func TestBytesBufferPool_Put(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		fields args
	}{
		{name: "size: 4096", fields: args{size: 4096}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &BytesBufferPool{
				pool: NewByteArrayPool(tt.fields.size, 0),
			}
			got := p.Get()
			got.WriteString("this is example.")
			p.Put(got)
			got = p.Get()

			if c := cap(got.Bytes()); c != tt.fields.size {
				t.Errorf("cap(bufpool.Get()) = %d, want %d", c, tt.fields.size)
			}
			if l := len(got.Bytes()); l != 0 {
				t.Errorf("len(bufpool.Get()) = %d, want 0", l)
			}
		})
	}
}

func TestNewBytesBufferPool(t *testing.T) {
	type args struct {
		size        int
		preallocate int
	}
	tests := []struct {
		name string
		args args
		want *BytesBufferPool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBytesBufferPool(tt.args.size, tt.args.preallocate)

			if c := got.pool.size; c != tt.args.size {
				t.Errorf("bufpool.size = %d, want %d", c, tt.args.size)
			}

			if c := got.pool.preallocate; c != tt.args.preallocate {
				t.Errorf("bufpool.preallocate = %d, want %d", c, tt.args.preallocate)
			}
		})
	}
}
