package ex6_test

import (
	"testing"
	"unicode"
)

func CollapseSpace(b []byte) []byte {
	running := false
	p := 0
	for _, v := range string(b) {
		if unicode.IsSpace(v) {
			if running {
				continue
			} else {
				running = true
				b[p] = ' '
				p++
			}
		} else {
			running = false
			// ?? is this the simplest way to get bytes from a rune?
			for _, c := range []byte(string(v)) {
				b[p] = c
				p++
			}
		}
	}
	return b[:p]
}

func TestCollapseSpace(t *testing.T) {
	b := CollapseSpace([]byte("\t\n\v\f\r "))
	if l := len(b); l != 1 {
		t.Errorf("expected len == 1, got len == %d, %s", l, b)
	}

	b = CollapseSpace([]byte("øHellø\t\n\v\f\rWørld™\t\n"))
	if s := string(b); s != "øHellø Wørld™ " {
		t.Errorf("expected 'øHellø Wørld™, got %s", s)
	}
}
