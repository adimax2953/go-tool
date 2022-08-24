package randtool

import (
	"math/rand"

	gotool "github.com/adimax2953/go-tool"
	"github.com/shopspring/decimal"
)

const (
	NMaxHit     = 5
	NEnlarge    = 300000
	BaseEnlarge = 100
)

var (
	seed5489 = int64(5489)
	gen5489  = uint64(14514284786278117030)
	orgorand = New(rand.NewSource(5489))
)

// Lottery - 長度4的陣列
func Lottery(values []int64) int {
	if len(values) != 4 {
		return 0
	}
	// Calculate Weight -
	var calcWeight [NMaxHit][2]int64
	var rtp [NMaxHit]int64 = [NMaxHit]int64{0, 100, 100, 100, 100}
	var rtpfix [NMaxHit]int64 = ConvertRTPFix([]int64{values[0], values[1], values[2], values[3]})
	var paytable [NMaxHit]int64 = [NMaxHit]int64{0, 1, 1, 1, 1}
	var paytablefix [NMaxHit]int64 = [NMaxHit]int64{0, rtpfix[1], rtpfix[2], rtpfix[3], rtpfix[4]}

	for idx := 1; idx < NMaxHit; idx++ {

		calcWeight[idx][1] = int64(NEnlarge * (rtp[idx] - rtpfix[idx]) / (paytable[idx] + paytablefix[idx]) / 1)
		if calcWeight[idx][1] > (BaseEnlarge * NEnlarge) {
			calcWeight[idx][1] = (BaseEnlarge * NEnlarge)
		}
		calcWeight[idx][0] = (BaseEnlarge * NEnlarge) - calcWeight[idx][1]
	}
	return OpenPoint(calcWeight)
}

func OpenPoint(calcWeight [NMaxHit][2]int64) int {

	for {
		rl := GetRandom(8)
		switch rl {
		case 0:
			// Decide Open or Not - by1234
			for idx := 1; idx < NMaxHit; idx++ {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			break
		case 1:
			// Decide Open or Not - by1234
			for idx := 2; idx < NMaxHit; idx++ {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[1][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 1
			}
			break
		case 2:
			// Decide Open or Not - by1234
			for idx := 3; idx < NMaxHit; idx++ {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[1][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 1
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[2][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 2
			}
			break
		case 3:
			// Decide Open or Not - by1234
			for idx := 4; idx < NMaxHit; idx++ {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[1][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 1
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[2][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 2
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[3][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 3
			}
			break
		case 4:
			// Decide Open or Not - by4321
			for idx := NMaxHit - 1; idx > 0; idx-- {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			break
		case 5:
			// Decide Open or Not - by4321
			for idx := NMaxHit - 2; idx > 0; idx-- {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[4][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 4
			}
			break
		case 6:
			// Decide Open or Not - by4321
			for idx := NMaxHit - 3; idx > 0; idx-- {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[4][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 4
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[3][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 3
			}
			break
		case 7:
			// Decide Open or Not - by4321
			for idx := NMaxHit - 4; idx > 0; idx-- {
				var weightArray []int32
				for i := 0; i < len(calcWeight[i]); i++ {
					weightArray = append(weightArray, int32(calcWeight[idx][i]))
				}
				if GenRandArray(weightArray, 2) != 0 {
					return idx
				}
			}
			var weightArray []int32
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[4][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 4
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[3][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 3
			}
			for i := 0; i < len(calcWeight[i]); i++ {
				weightArray = append(weightArray, int32(calcWeight[2][i]))
			}
			if GenRandArray(weightArray, 2) != 0 {
				return 2
			}
			break
		}
	}
}

func ConvertRTPFix(value []int64) [NMaxHit]int64 {
	if len(value) != 4 {
		return [NMaxHit]int64{0, 0, 0, 0, 0}
	}
	return minNum(value)
}

func minNum(value []int64) (rtpfix [NMaxHit]int64) {
	if len(value) != 4 {
		return [NMaxHit]int64{0, 25, 25, 25, 25}
	}
	var max int64
	for _, val := range value {
		max += val
	}

	rtpfix[1] = gotool.Str2int64(decimal.NewFromInt(value[0]).Div(decimal.NewFromInt(max)).Mul(decimal.NewFromInt(100)).Floor().String())
	rtpfix[2] = gotool.Str2int64(decimal.NewFromInt(value[1]).Div(decimal.NewFromInt(max)).Mul(decimal.NewFromInt(100)).Floor().String())
	rtpfix[3] = gotool.Str2int64(decimal.NewFromInt(value[2]).Div(decimal.NewFromInt(max)).Mul(decimal.NewFromInt(100)).Floor().String())
	rtpfix[4] = gotool.Str2int64(decimal.NewFromInt(value[3]).Div(decimal.NewFromInt(max)).Mul(decimal.NewFromInt(100)).Floor().String())

	return rtpfix
}
