// Copyright 2019 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufpool

import (
	"testing"
)

func Test_ByteArrayPool_Get(t *testing.T) {
	type fields struct {
		size        int
		preallocate int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "size: 4096, preallocate: 0", fields: fields{size: 4096, preallocate: 0}},
		{name: "size: 4096, preallocate: 4", fields: fields{size: 4096, preallocate: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ByteArrayPool{
				size:        tt.fields.size,
				preallocate: tt.fields.preallocate,
			}
			got := p.Get()
			if c := cap(got); c != tt.fields.size {
				t.Errorf("cap(bufpool.Get()) = %d, want %d", c, tt.fields.size)
			}
			if l := len(got); l != 0 {
				t.Errorf("len(bufpool.Get()) = %d, want 0", l)
			}
		})
	}
}

func Test_ByteArrayPool_Put(t *testing.T) {
	type fields struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "size: 4096", fields: fields{size: 4096}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ByteArrayPool{
				size: tt.fields.size,
			}
			got := p.Get()
			got = append(got, []byte("this is example.")...)
			p.Put(got)
			got = p.Get()

			if c := cap(got); c != tt.fields.size {
				t.Errorf("cap(bufpool.Get()) = %d, want %d", c, tt.fields.size)
			}
			if l := len(got); l != 0 {
				t.Errorf("len(bufpool.Get()) = %d, want 0", l)
			}
		})
	}
}

func TestNewByteArrayPool(t *testing.T) {
	type args struct {
		size        int
		preallocate int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "size: 4096, preallocate: 0", args: args{size: 4096, preallocate: 0}},
		{name: "size: 4096, preallocate: 4", args: args{size: 4096, preallocate: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewByteArrayPool(tt.args.size, tt.args.preallocate)

			if c := got.size; c != tt.args.size {
				t.Errorf("bufpool.size = %d, want %d", c, tt.args.size)
			}

			if c := got.preallocate; c != tt.args.preallocate {
				t.Errorf("bufpool.preallocate = %d, want %d", c, tt.args.preallocate)
			}
		})
	}
}
