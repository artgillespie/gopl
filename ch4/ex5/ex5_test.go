package ex5_test

import "testing"

func Dedup(strings []string) []string {
	i := 1
	for j := 1; j < len(strings); j++ {
		if strings[j] == strings[j-1] {
			continue
		}
		strings[i] = strings[j]
		i++
	}
	return strings[:i]
}

func TestDeDup(t *testing.T) {
	var s = []string{"one", "two", "two", "three"}
	s2 := Dedup(s)
	t.Logf("s2: %v, s: %v", s2, s)
}
