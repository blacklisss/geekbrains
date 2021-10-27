package fibonacci_test

import (
	. "geekbrains/go/Lesson_10/fibonacci"
	"geekbrains/go/Lesson_10/mycache"
	"testing"
)

var mapCache = mycache.NewCache()

func BenchmarkFibStruct_CalculateWithoutCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fb.Calculate(40, false)
	}
}

func BenchmarkFibStruct_CalculateWithCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fb.Calculate(40, true)
	}
}

func init() {
	Fb.Cache = mapCache
}
