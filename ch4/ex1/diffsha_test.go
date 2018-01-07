package diffsha_test

import "testing"
import "crypto/sha256"

// DifSHA256 returns the number of differing bits between two SHA256 digests.
func DiffSHA256(x, y [32]uint8) int {
	var c int
	for i, _ := range x {
		d := x[i] ^ y[i]
		c += Count(d)
	}
	return c
}

// Count counts the number of set bits in a byte
func Count(b uint8) int {
	var count int
	for i := uint(0); i < 8; i++ {
		count += int(b >> i & 0x1)
	}
	return count
}

func TestCount(t *testing.T) {
	if c := Count(1); c != 1 {
		t.Fatalf("Count(1) expected 1, got %d", c)
	}
	if c := Count(3); c != 2 {
		t.Fatalf("Count(3) expected 2, got %d", c)
	}
	if c := Count(255); c != 8 {
		t.Fatalf("Count(255) expected 8, got %d", c)
	}
}

func expectDigestDiff(t *testing.T, a, b string, expected int) {
	x := sha256.Sum256([]byte(a))
	y := sha256.Sum256([]byte(b))
	if d := DiffSHA256(x, y); d != expected {
		t.Fatalf("Expected bit difference to be %d, got %d", expected, d)
	}
}

func TestDiffSHA256(t *testing.T) {
	expectDigestDiff(t, "x", "X", 125)
	expectDigestDiff(t, "Hello, World!", "Hello, World!", 0)
}
