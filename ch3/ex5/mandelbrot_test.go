package main_test

import "testing"

func TestColorScale(t *testing.T) {
	const mid uint16 = 255/2 + 255
	if mid != 382 {
		t.Fatalf("Expected 383, got %d", mid)
	}
}
