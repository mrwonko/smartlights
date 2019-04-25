package benchmark

import (
	"testing"
	"time"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Millisecond)
	}
}
