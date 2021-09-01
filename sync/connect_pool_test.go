//@Author: zhuo.chen(c-Erdong@gmail.com)
//@Date:   2021/8/31 4:14 PM
//Copyright 2021 www.ucloud.cn

package main

import (
	"sync"
	"testing"
)

var bytePool = sync.Pool{
	New: newPool,
}

func newPool() interface{} {
	b := make([]byte, 1024)
	return &b
}

func BenchmarkAlloc(b *testing.B) {
	for i:=0; i<b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

func BenchmarkPool(b *testing.B) {
	for i:=0; i<b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}
