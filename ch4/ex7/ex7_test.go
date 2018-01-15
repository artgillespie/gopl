package ex7_test

import (
	"log"
	"testing"
	"unicode/utf8"
)

// Reverse ex4.7 — Modify reverse to reverse the characters of a []byte slice that
// represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
func Reverse(b []byte) {
	// DecodeRune -> use width to copy to b2[j-w:j]
	b2 := make([]byte, len(b))
	log.Printf("b2: %d\n", len(b2))
	for i, w := 0, 0; i < len(b); i += w {
		_, width := utf8.DecodeRune(b[i:])
		y := len(b2) - (width) - i
		z := y + (width)
		log.Printf("y: %d, %d, %d, %d\n", y, z, i, width)
		copy(b2[y:z], b[i:width])
		w = width
	}
	copy(b, b2)
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

func TestReverse(t *testing.T) {
	b := []byte("ø∫")
	t.Logf("%s", b)
	Reverse(b)
	t.Logf("%s", b)
	//t.Logf("Rune: %v", RuneAtIndex([]byte("ø™ø"), 1))
}
