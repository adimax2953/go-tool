package randtool

import (
	"math/rand"
	"sync"
)

type Rand struct {
	lk   sync.Mutex
	rand *rand.Rand
}

// New returns a new gosfmt Rand that uses random values from src
// to generate other random values.
func New(source rand.Source) *Rand {
	return &Rand{
		rand: rand.New(source),
	}
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *Rand) Seed(seed int64) {
	r.lk.Lock()
	r.rand.Seed(seed)
	r.lk.Unlock()
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 {
	r.lk.Lock()
	val := r.rand.Int63()
	r.lk.Unlock()
	return val
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 {
	r.lk.Lock()
	val := r.rand.Uint32()
	r.lk.Unlock()
	return val
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64 {
	r.lk.Lock()
	val := r.rand.Uint64()
	r.lk.Unlock()
	return val
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 {
	r.lk.Lock()
	val := r.rand.Int31()
	r.lk.Unlock()
	return val
}

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	r.lk.Lock()
	val := r.rand.Int()
	r.lk.Unlock()
	return val
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	r.lk.Lock()
	val := r.rand.Int63n(n)
	r.lk.Unlock()
	return val
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	r.lk.Lock()
	val := r.rand.Int31n(n)
	r.lk.Unlock()
	return val
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	r.lk.Lock()
	val := r.rand.Intn(n)
	r.lk.Unlock()
	return val
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 {
	r.lk.Lock()
	val := r.rand.Float64()
	r.lk.Unlock()
	return val
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 {
	r.lk.Lock()
	val := r.rand.Float32()
	r.lk.Unlock()
	return val
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int {
	r.lk.Lock()
	val := r.rand.Perm(n)
	r.lk.Unlock()
	return val
}

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int)) {
	r.lk.Lock()
	r.rand.Shuffle(n, swap)
	r.lk.Unlock()
}

// Read generates len(p) random bytes and writes them into p. It
// always returns len(p) and a nil error.
// Read should not be called concurrently with any other Rand method.
func (r *Rand) Read(p []byte) (n int, err error) {
	r.lk.Lock()
	n, err = r.rand.Read(p)
	r.lk.Unlock()
	return n, err
}

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default Source.
// To produce a different normal distribution, callers can
// adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func (r *Rand) NormFloat64() float64 {
	r.lk.Lock()
	nf := r.rand.NormFloat64()
	r.lk.Unlock()
	return nf
}

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default Source.
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter
func (r *Rand) ExpFloat64() float64 {
	r.lk.Lock()
	ef := r.rand.ExpFloat64()
	r.lk.Unlock()
	return ef
}
