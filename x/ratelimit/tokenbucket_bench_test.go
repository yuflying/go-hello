package ratelimit

import (
	"testing"
	"time"
)

func BenchmarkLimit(b *testing.B) {
	rl := New(10, time.Second)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rl.Limit()
		}
	})
}
