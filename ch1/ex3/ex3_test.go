package ex3_test

import (
	"strings"
	"testing"
)

func echo1(args []string) string {
	var s, sep string
	for _, a := range args {
		s += sep + a
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	return strings.Join(args, " ")
}

var data = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen"}
var data2 = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty-one", "twenty-two", "twenty-three", "twenty-four", "twenty-five", "twenty-six", "twenty-seven", "twenty-eight", "twenty-nine", "thirty"}

// BenchmarkEcho1 benchmarks the naive implementation of echo
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(data)
	}
}
func BenchmarkEcho1_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(data2)
	}
}

// BenchmarkEcho2 benchmarks the strings.Join implementation of echo
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(data)
	}
}

func BenchmarkEcho2_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(data2)
	}
}
