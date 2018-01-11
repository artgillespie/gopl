package futures_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	// basically hacking at https://medium.com/strava-engineering/futures-promises-in-the-land-of-golang-1453f4807945
	afunc := func(s int) func() (int, error) {
		return func() (int, error) {
			time.Sleep(time.Duration(s) * time.Millisecond)
			return 0, fmt.Errorf("WE EXPECTED THIS ERROR AFTER %d MILLISECONDS", s)
		}
	}

	future := func(f func() (int, error)) func() (int, error) {
		c := make(chan int)
		cerr := make(chan error)

		go func() {
			i, err := f()
			if err != nil {
				cerr <- err
				return
			}
			c <- i
		}()

		return func() (int, error) {
			select {
			case i := <-c:
				return i, nil
			case err := <-cerr:
				return 0, err
			}
		}
	}

	r := future(afunc(2000))
	log.Printf("future fired...\n")
	v, err := r()
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if v != 2000 {
		t.Errorf("Expected 1000, got %d", v)
	}
}

func latencyEcho(n float64) (float64, error) {
	time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
	return n, fmt.Errorf("A TOTALLY UNEXPECTED ERROR!")
}

func TestFanOutIn(t *testing.T) {
	// this ensures the latency of latencyEcho isn't stable between test runs
	rand.Seed(time.Now().Unix())
	// the values we'll pass to our 'promises'
	var values = []float64{1.0, 1.5, 2.0, 3.0}

	// the channel the promises will communicate their response back on
	c := make(chan float64)
	cerr := make(chan error)
	// for each value, fire off a goroutine which runs in parallel
	for _, v := range values {
		go func(n float64) {
			v, err := latencyEcho(n)
			if err != nil {
				cerr <- err
				return
			}
			c <- v
		}(v)
	}

	// Fan-in: this next bit is like Promise.all with a timeout
	// see https://talks.golang.org/2012/concurrency.slide#27
	fanIn := func(c <-chan float64, n int) ([]float64, error) {
		acc := make(chan []float64)
		var a []float64
		go func() {
			for i := 0; i < n; i++ {
				a = append(a, <-c)
			}
			acc <- a
		}()
		select {
		case results := <-acc:
			return results, nil
		case err := <-cerr:
			return nil, err
		case <-time.After(900 * time.Millisecond):
			return nil, fmt.Errorf("Timed out")
		}
	}

	results, err := fanIn(c, len(values))
	if err != nil {
		t.Logf("Error: %v", err)
		return
	}
	t.Logf("Got Results: %v", results)

}
