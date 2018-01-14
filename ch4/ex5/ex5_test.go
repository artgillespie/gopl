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
	var tests = []struct {
		a []string
		b []string
	}{
		{
			[]string{"one", "two", "two", "three"},
			[]string{"one", "two", "three"},
		},
	}
	for _, test := range tests {
		s2 := Dedup(test.a)
		if l := len(s2); l != len(test.b) {
			t.Errorf("Expected %d to equal %d", l, len(test.b))
		}
		for i, v := range s2 {
			if v != s2[i] {
				t.Errorf("Expected %s to equal %s", v, s2[i])
			}
		}
	}
}
