package extra_test

import (
	"fmt"
	"testing"
)

type IntQueue []int

func (q *IntQueue) Push(i int) {
	*q = append(*q, i)
}

func (q *IntQueue) Pop() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return r, nil
}

func TestIntQueue(t *testing.T) {
	var q IntQueue
	q.Push(0)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	verify := func(q *IntQueue, val int) {
		if r, err := (*q).Pop(); err != nil {
			t.Errorf("unexpected error: %v", err)
		} else if r != val {
			t.Errorf("expected %d, got %d", val, r)
		}
	}

	verify(&q, 0)
	verify(&q, 1)
	verify(&q, 2)
	verify(&q, 3)
	_, err := q.Pop()
	if err == nil {
		t.Errorf("expected an error")
	}

}
