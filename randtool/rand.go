package randtool

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"

	"os"
	"sync"
	"time"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/seehuhn/mt19937"
)

var rngPool sync.Pool

var mt19937Rand = New(mt19937.New())

func init() {
	b := new(big.Int).SetUint64(uint64(time.Now().UTC().UnixNano() / int64(os.Getpid())))
	sd, _ := crand.Int(crand.Reader, b)
	x := sd.Uint64() + 0x9E3779B97F4A7C15
	x ^= x >> 30 * 0xBF58476D1CE4E5B9
	x ^= x << 27 * 0x94D049BB133111EB
	x ^= x >> 31
	seed := int64(x)

	mt19937Rand.Seed(seed)
	rand.Seed(seed)
	LogTool.LogSystem("init Rng")
}

// Uint32 - returns pseudorandom uint32.
//
// It is safe calling this function from concurrent goroutines.
func Uint32() uint32 {
	return mt19937Rand.Uint32()
}

// Uint32n - safe
func Uint32n(maxN uint32) uint32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

// GetRandom - safe
func GetRandom(maxN int32) int32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return int32((uint64(x) * uint64(maxN)) >> 32)
}

// GenRandArray - safe
func GenRandArray(weightArray []int32, arraySizze int32) uint32 {
	var resultNum uint32
	var sumWeight uint32
	var sumArray []uint32
	sumArray = make([]uint32, arraySizze)
	var i int32

	for i = 0; i < arraySizze; i++ {
		sumWeight += uint32(weightArray[i])
		sumArray[i] = sumWeight
	}

	var randNum uint32
	randNum = Uint32n(sumWeight)

	for i = 0; i < arraySizze; i++ {
		if randNum < sumArray[i] {
			resultNum = uint32(i)
			break
		}
	}

	return resultNum
}

// RNG is a pseudorandom number generator.
//
// It is unsafe to call RNG methods from concurrent goroutines.
type RNG struct {
	state uint32
}

// GetRandom -
func (r *RNG) GetRandom(maxN int32) int32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return int32((uint64(x) * uint64(maxN)) >> 32)
}

// GenRandArray -
func (r *RNG) GenRandArray(weightArray []int32, arraySizze int32) uint32 {
	var resultNum uint32
	var sumWeight uint32
	var sumArray []uint32
	sumArray = make([]uint32, arraySizze)
	var i int32

	for i = 0; i < arraySizze; i++ {
		sumWeight += uint32(weightArray[i])
		sumArray[i] = sumWeight
	}

	var randNum uint32
	randNum = r.Uint32n(sumWeight)

	for i = 0; i < arraySizze; i++ {
		if randNum < sumArray[i] {
			resultNum = uint32(i)
			break
		}
	}

	return resultNum
}

// Uint32 returns pseudorandom uint32.
//
// It is unsafe to call this method from concurrent goroutines.
func (r *RNG) Uint32() uint32 {
	if r.state == 0 {
		r.state = getRandomUint32()
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	x := r.state
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.state = x
	return x
}

// Uint32n returns pseudorandom uint32 in the range [0..maxN).
//
// It is unsafe to call this method from concurrent goroutines.
func (r *RNG) Uint32n(maxN uint32) uint32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

func getRandomUint32() uint32 {
	// x := time.Now().UnixNano()
	x := uint64(time.Now().UTC().UnixNano() / int64(os.Getpid()))
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	return uint32((x >> 32) ^ x)
}

// Shuffle -打亂陣列
func Shuffle[T NonNegative_Integer](nums []T) []T {
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		nums[last], nums[idx] = nums[idx], nums[last]
	}
	return nums
}
