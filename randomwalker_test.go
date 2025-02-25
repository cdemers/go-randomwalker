package randomwalker

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

type randSim struct{}

func (r randSim) Int63() int64 {
	return 0
}

func (r randSim) Seed(seed int64) {}

func TestNewRandomWalker(t *testing.T) {
	rw := NewRandomWalker(10, 5, 15, 0)
	value := rw.Step()
	if value != 10 {
		t.Errorf("Expected %0.2f, got %0.2f", 10.0, value)
	}

	rw = NewRandomWalker(10, 5, 15, 10)
	for q := 0; q < 1000; q++ {
		value = rw.Step()
		if value > 15 || value < 5 {
			t.Errorf("Expected value between %0.2f and %0.2f, got %0.2f", 5.0, 15.0, value)
		}
	}

	var source rand.Source = randSim{}

	S := source.(rand.Source)
	rw = NewRandomWalkerWithRandSource(10, 5, 15, 0.25, &S)

	value = rw.Step()
	if value != 7.5 {
		t.Errorf("Expected value to be %0.5f, got %0.5f", 7.5, value)
	}
	value = rw.Step()
	if value != 5.625 {
		t.Errorf("Expected value to be %0.5f, got %0.5f", 5.625, value)
	}
	value = rw.Step()
	if value != 5.0 {
		t.Errorf("Expected value to be %0.5f, got %0.5f", 5.0, value)
	}

}

func ExampleRandomWalker_Step() {
	rw := NewRandomWalker(10, 5, 15, 0)
	value := rw.Step()
	fmt.Println(value)
	// Output: 10
}

func TestRandomWalker_Concurrency(t *testing.T) {
	rw := NewRandomWalker(10, 5, 15, 0.1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				rw.Step()
			}
		}()
	}
	wg.Wait()
}
