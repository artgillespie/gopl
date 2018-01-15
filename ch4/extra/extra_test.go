package extra_test

import (
	"fmt"
	"testing"
)

type IntStack []int

func (s *IntStack) Push(i int) {
	*s = append(*s, i)
}

func (s *IntStack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("IntStack length == 0")
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r, nil
}

func TestIntStack(t *testing.T) {
	var s IntStack
	s.Push(0)
	s.Push(1)
	s.Push(2)
	if v, err := s.Pop(); err != nil {
		t.Errorf("Pop() error %v", err)
	} else if v != 2 {
		t.Errorf("Expected 2, got %d", v)
	}
	if v, err := s.Pop(); err != nil {
		t.Errorf("Pop() error %v", err)
	} else if v != 1 {
		t.Errorf("Expected 1, got %d", v)
	}
	if v, err := s.Pop(); err != nil {
		t.Errorf("Pop() error %v", err)
	} else if v != 0 {
		t.Errorf("Expected 0, got %d, %v", v, s)
	}
	if l := len(s); l != 0 {
		t.Errorf("Expected len == 0, got %d, %v", l, s)
	}
}
