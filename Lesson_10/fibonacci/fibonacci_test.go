package fibonacci

import (
	"geekbrains/go/Lesson_10/mycache"
	"testing"
)

var mapCache = mycache.NewCache()
var KnownFibList = []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}

func TestFibZero(t *testing.T) {
	if Fb.Calculate(0, false) != 0 {
		t.Errorf("Fibonacci(0) expected to be zero")
	}

}

func TestFibOne(t *testing.T) {
	if Fb.Calculate(1, false) != 1 {
		t.Errorf("Fibonacci(1) expected to be one")
	}

}

func TestFibMultiWithoutCache(t *testing.T) {

	for i, expected := range KnownFibList {
		if received := Fb.Calculate(int64(i), false); received != expected {
			t.Errorf("Fibonacci(%d) expected %d, but received %d", i, expected, received)
		}
	}
}

func TestFibMultiWithCache(t *testing.T) {

	for i, expected := range KnownFibList {
		if received := Fb.Calculate(int64(i), true); received != expected {
			t.Errorf("Fibonacci(%d) expected %d, but received %d", i, expected, received)
		}
	}
}

func init() {
	Fb.Cache = mapCache
}
