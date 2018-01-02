package popcount_test

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountA(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountC(x uint64) int {
	var c int
	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return c
}

func assertEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("Expected %d to equal %d", a, b)
	}
}

func TestPopCountC(t *testing.T) {
	assertEqual(t, PopCountC(0), 0)
	assertEqual(t, PopCountC(255), 8)
	assertEqual(t, PopCountC(511), 9)
	assertEqual(t, PopCountC(513), 2)
	assertEqual(t, PopCountC(1024), 1)
}

func BenchmarkPopCountC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountC(124567)
	}
}

func BenchmarkPopCountA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountA(124567)
	}
}
