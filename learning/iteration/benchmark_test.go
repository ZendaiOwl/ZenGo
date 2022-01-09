package iteration

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	var repeated string = "a"
	for i := 0; i < b.N; i++ {
		Repeat(repeated, i)
	}
}
