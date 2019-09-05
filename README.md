# Simple Random Walker

**Simple random walker to be used for non cryptographic simulations, like
natural phenomenon models or finance**

Package randomwalker provides a parametric random walk generator. You can
instantiate a random walker that will start at a specific point, that will stay
between an upper and lower boundary, that will vary on each step by a specified
maximum magnitude.

Also, if you expect to use this walker for financial models, this walker is
going to give you something close to a Gaussian random walk, and not a Brownian
random walk.

Calling NewRandomWalker() returns a Gaussian like random walk generator.  Once
created, you can call Step() on this object to get the first and then the
subsequent random walk value.

It's signature is `NewRandomWalker(origin, min, max, maxDynPcnt float32)`.

- The _origin_ parameter sets the initial walk value, it's starting point.
- The _min_ parameter sets a limit on how low can the walk go.
- The _max_ parameter sets a limit on how high the walk can go.
- The _maxDynPcnt_ parameter sets the limit on how much can a walk step vary from
  it's previous value. Setting this value at zero will prevent the walk from
  varying at all, leaving it constantly at it's initial value. Setting this
  value at 1.0 will allow the walk step to vary from -50% to +50% of it's
  previous value (the difference between -50% and +50% totaling 1.0). And
  setting this value to more than 1.0 is permitted, for example setting it to 4.0
  will allow a step to vary between -200% and +200% of it's previous value.

Calling `Step()` afterward will return a float32 value.

Using NewRandomWalkerWithRandSource() also returns a random walk generator but
that here you can provide your own random source object if needed.

**Example**

```go
var origin, minimumWalkValue, maximumWalkValue, maximumVariationPercentage float32

origin = 50
minimumWalkValue = -25
maximumWalkValue = 75
maximumVariationPercentage = 0.01

rw = NewRandomWalker(origin, minimumWalkValue, maximumWalkValue, maximumVariationPercentage)
value = rw.Step()
// _value_ will be somewhere between (50 + 1%) and (50 - 1%).
value = rw.Step()
// _value_ will now be it's previous value plus or minus 1%.
```
