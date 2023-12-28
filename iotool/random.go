package iotool

import (
	cr "crypto/rand"
	"math"
	"math/big"
	mr "math/rand"
	"reflect"

	LogTool "github.com/adimax2953/log-tool"
)

var seed *big.Int

type RandomValType interface {
	int | int32 | int64
}

func init() {
	var err error
	seed, err = cr.Int(cr.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		LogTool.LogWarningf("RandomSeedInit", "Random Seed error, %v", err)
	} else {
		LogTool.LogInfo("RandomSeedInit", "Random seed created.", seed)
	}
	//set seed
	mr.NewSource(seed.Int64())
}

// RandIntTn 隨機產生0到n-1之亂數(接受型別[int | int32 | int64])
func RandIntTn[T RandomValType](n T) T {
	if n <= 0 {
		return 0
	}
	switch reflect.TypeOf(n).Kind() {
	case reflect.Int:
		return T(mr.Intn(int(n)))
	case reflect.Int32:
		return T(mr.Int31n(int32(n)))
	case reflect.Int64:
		return T(mr.Int63n(int64(n)))
	}
	return T(mr.Int63n((int64(n))))
}

// IsInProbability 是否有中(n分之1)
func IsInProbability[T RandomValType](n T) bool {
	if n <= 0 {
		return false
	}
	return RandIntTn[T](n) == 0
}

// RandSliceIndex 對slice亂數取index
func RandSliceIndex(inSlice interface{}) (int, bool) {
	switch reflect.TypeOf(inSlice).Kind() {
	case reflect.Slice:
		length := reflect.ValueOf(inSlice).Len()
		return RandIntTn[int](length), true
	default:
		return -1, false
	}
}

// ShuffleSlice 隨機打亂slice裡的順序
func ShuffleSlice(inSlice interface{}) {
	swap := reflect.Swapper(inSlice)
	length := reflect.ValueOf(inSlice).Len()
	for i := length - 1; i > 0; i-- {
		j := mr.Intn(i + 1)
		swap(i, j)
	}
}
