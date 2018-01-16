package ex7_test

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

// Reverse ex4.7 — Modify reverse to reverse the characters of a []byte slice that
// represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
//                   Reverse Runes     Reverse entire buffer
// [[123], [456]] -> [[321], [654]] -> [456, 123]
func ReverseUTF8(b []byte) {
	for i, w := 0, 0; i < len(b); i += w {
		_, width := utf8.DecodeRune(b[i:])
		Reverse(b[i : i+width])
		w = width
	}
	Reverse(b)
}

func Reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}
}

func TestReverse(t *testing.T) {
	b := []byte("Hello")
	Reverse(b)
	if bytes.Compare(b, []byte{'o', 'l', 'l', 'e', 'H'}) != 0 {
		t.Errorf("Expected 'olleH', got %s", b)
	}
}

func RuneAtIndex(b []byte, idx int) rune {
	for i, w := 0, 0; i < len(b); i++ {
		r, width := utf8.DecodeRune(b[w:])
		if i == idx {
			return r
		}
		w += width
	}
	return ' '
}

func TestReverseUTF8(t *testing.T) {
	b := []byte("ø∫™ªº¡")
	ReverseUTF8(b)
	e := "¡ºª™∫ø"
	if bytes.Compare([]byte(e), b) != 0 {
		t.Errorf("Expected %s, got %s", e, b)
	}
}
