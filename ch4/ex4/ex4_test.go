package ex4_test

import (
	"testing"

	"github.com/artgillespie/gopl/ch4/ex4"
)

func assertLen(t *testing.T, s []int, e int) {
	if l := len(s); l != e {
		t.Fatalf("Assert Len: expected %d, got %d", e, l)
	}
}

func assertNil(t *testing.T, s []int) {
	if s != nil {
		t.Fatalf("Assert Nil: expected nil, got %v", s)
	}
}

func TestSliceAppend(t *testing.T) {
	a := []int{} // empty slice literal
	var b []int  // nil slice

	t.Logf("%T, %T", a, b)

	assertLen(t, a, 0)
	assertLen(t, b, 0)
	assertNil(t, b)

	b = append(b, 2)
	a = append(a, 2)
	assertLen(t, a, 1)
	assertLen(t, b, 1)
}

func TestSliceIndex(t *testing.T) {
	a := []int{1}
	var b []int = []int{}

	b[0] = 1
	a[0] = 1
}

func equalArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func assertEqualArray(t *testing.T, e []int, r []int) {
	if !equalArray(e, r) {
		t.Fatalf("Equal Array: Expected %v, got %v", e, r)
	}
}

func TestRotate(t *testing.T) {
	var tests = []struct {
		a []int
		e []int
		n int
	}{
		{
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 0,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}, 1,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{4, 5, 1, 2, 3}, 2,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 1}, -1,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, -2,
		},
		{
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, -5,
		},
	}

	for _, test := range tests {
		ex4.Rotate(test.a, test.n)
		assertEqualArray(t, test.e, test.a)
	}
}
