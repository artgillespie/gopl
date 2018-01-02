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

// Book implementation
// 3.87 ns/op
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

// Ex 2.3
// 17.1 ns/op
func PopCountB(x uint64) int {
	var c int
	var i uint
	for i = 0; i < 8; i++ {
		c += int(pc[byte(x>>(i*8))])
	}
	return c
}

// Ex 2.4
// 35.4 ns/op
func PopCountC(x uint64) int {
	var c int
	for i := 0; i < 64; i++ {
		c += int(x & 1)
		x >>= 1
	}
	return c
}

// Ex 2.5
// 5.6 ns/op
func PopCountD(x uint64) int {
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

func validate(t *testing.T, f func(uint64) int) {
	for _, v := range testValues {
		assertEqual(t, f(v), PopCountA(v))
	}
}

var testValues = []uint64{255, 256, 511, 513, 255123, 1, 1000000}

func TestPopCountB(t *testing.T) {
	validate(t, PopCountB)
}

func TestPopCountC(t *testing.T) {
	validate(t, PopCountC)
}

func TestPopCountD(t *testing.T) {
	validate(t, PopCountD)
}

func runIterations(f func(uint64) int, n int) {
	for i := 0; i < n; i++ {
		f(124567)
	}
}

func BenchmarkPopCountA(b *testing.B) {
	runIterations(PopCountA, b.N)
}

func BenchmarkPopCountB(b *testing.B) {
	runIterations(PopCountB, b.N)
}

func BenchmarkPopCountC(b *testing.B) {
	runIterations(PopCountC, b.N)
}

func BenchmarkPopCountD(b *testing.B) {
	runIterations(PopCountD, b.N)
}
