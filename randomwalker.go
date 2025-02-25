/*
Package randomwalker provides a parametric random walk generator. You can
instentiate a random walker that will start at a specific point, that will stay
between an upper and lower boundary, that will vary on each step by a specified
maximum magnetude.

Also, if you expect to use this walker for financial models, this walker is
going to give you something close to a Gaussian random walk, and not a Brownian
random walk.
*/
package randomwalker

import (
	"math/rand"
	"sync"
	"time"
)

// randomer abstracts the random number generator.  We would not need this if
// it was not for the tests, which need a "replayable" random source, so we
// use this inteface to be able to stub the random generator in the tests.
type randomer interface {
	Float32() float32
}

// RandomWalker represents one instance of a random walk generator.
type RandomWalker struct {
	mu         sync.Mutex
	current    float32
	min        float32
	max        float32
	maxDynPcnt float32
	random     randomer
}

// NewRandomWalker returns a random walk generator.  Once created, you can call
// Step() on this object to get a random walk value.
func NewRandomWalker(origin, min, max, maxDynPcnt float32) *RandomWalker {
	source := rand.NewSource(time.Now().UnixNano())
	return NewRandomWalkerWithRandSource(origin, min, max, maxDynPcnt, &source)
}

// NewRandomWalkerWithRandSource returns a random walk generator that will use
// the provided random source.  Specifying the random can help making sure that
// parallel processes will not produce similar random walks, or if you want to
// skew the random walk to better fit a certain natural phenomenon you are
// trying to simulate. Once created, you can call Step() on this object to get
// a random walk value.
func NewRandomWalkerWithRandSource(origin, min, max, maxDynPcnt float32, source *rand.Source) *RandomWalker {
	return &RandomWalker{
		current:    origin,
		min:        min,
		max:        max,
		maxDynPcnt: maxDynPcnt,
		random:     rand.New(*source),
	}
}

// Step does one step of a random walk and returns the resulting value.
func (rw *RandomWalker) Step() float32 {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	maxDynamic := rw.current * rw.maxDynPcnt

	// BUG(cdemers): The pseudorandom generator seems to be skewed on the
	// low side by about 3%, which is extremely strange.
	rw.current += (rw.random.Float32()*2 - 1) * maxDynamic

	if rw.current < rw.min {
		rw.current = rw.min
	}
	if rw.current > rw.max {
		rw.current = rw.max
	}

	return rw.current
}
