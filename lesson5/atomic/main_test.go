package main

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicAdd(b *testing.B) {
	a := uint64(0)
	for i := 0; i < b.N; i++ {
		atomic.AddUint64(&a, 1)
	}
}

func BenchmarkAdd(b *testing.B) {
	a := uint64(0)
	for i := 0; i < b.N; i++ {
		a += 1
	}
}
