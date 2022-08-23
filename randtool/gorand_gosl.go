package randtool

// Int63r generates pseudo random int64 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random int64
func (r *Rand) Int63r(low, high int64) int64 {
	return r.Int63()%(high-low+1) + low
}

// Int63s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Int63s(values []int64, low, high int64) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int63r(low, high)
	}
}

// Int63Shuffle - shuffles a slice of integers
func (r *Rand) Int63Shuffle(values []int64) {
	var tmp int64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Uint32 is int range generates pseudo random uint32 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random uint32
func (r *Rand) Uint32r(low, high uint32) uint32 {
	return r.Uint32()%(high-low+1) + low
}

// Uint32s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Uint32s(values []uint32, low, high uint32) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint32r(low, high)
	}
}

// Uint32Shuffle shuffles a slice of integers
func (r *Rand) Uint32Shuffle(values []uint32) {
	var tmp uint32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Uint64r generates pseudo random uint64 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random uint64
func (r *Rand) Uint64r(low, high uint64) uint64 {
	return r.Uint64()%(high-low+1) + low
}

// Uint64s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Uint64s(values []uint64, low, high uint64) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint64r(low, high)
	}
}

// Uint64Shuffle - shuffles a slice of integers
func (r *Rand) Uint64Shuffle(values []uint64) {
	var tmp uint64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Int31r is int range generates pseudo random int32 between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random int32
func (r *Rand) Int31r(low, high int32) int32 {
	return r.Int31()%(high-low+1) + low
}

// Int31s generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Int31s(values []int32, low, high int32) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int31r(low, high)
	}
}

// Int31Shuffle - shuffles a slice of integers
func (r *Rand) Int31Shuffle(values []int32) {
	var tmp int32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Intr is int range generates pseudo random integer between low and high.
//  Input:
//   low  -- lower limit
//   high -- upper limit
//  Output:
//   random integer
func (r *Rand) Intr(low, high int) int {
	return r.Int()%(high-low+1) + low
}

// Ints generates pseudo random integers between low and high.
//  Input:
//   low    -- lower limit
//   high   -- upper limit
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Ints(values []int, low, high int) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Intr(low, high)
	}
}

// IntShuffle shuffles a slice of integers
func (r *Rand) IntShuffle(values []int) {
	var j, tmp int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Float64r generates a pseudo random real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   random float64
func (r *Rand) Float64r(low, high float64) float64 {
	return low + (high-low)*r.Float64()
}

// Float64s generates pseudo random real numbers between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Float64s(values []float64, low, high float64) {
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float64()
	}
}

// Float64Shuffle shuffles a slice of float point numbers
func (r *Rand) Float64Shuffle(values []float64) {
	var tmp float64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Float32r generates a pseudo random real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   random float32
func (r *Rand) Float32r(low, high float32) float32 {
	return low + (high-low)*r.Float32()
}

// Float32s generates pseudo random real numbers between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit (closed)
//   high -- upper limit (open)
//  Output:
//   values -- slice to be filled with len(values) numbers
func (r *Rand) Float32s(values []float32, low, high float32) {
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float32()
	}
}

// Float32Shuffle shuffles a slice of float point numbers
func (r *Rand) Float32Shuffle(values []float32) {
	var tmp float32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// FlipCoin generates a Bernoulli variable; throw a coin with probability p
func (r *Rand) FlipCoin(p float64) bool {
	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	if r.Float64() <= p {
		return true
	}
	return false
}
