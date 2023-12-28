package iotool

import (
	"log"
)

type sliceType interface {
	string | int | int32 | int64
}

type ProbWeight struct {
	probSlice []int64
	total     int64
}

// CreateProb 建立一個權重結構
func CreateProb(inProb ...int64) ProbWeight {

	var prob ProbWeight

	prob.probSlice = make([]int64, len(inProb))

	for i := 0; i < len(inProb); i++ {
		prob.probSlice[i] = inProb[i]
		prob.total += inProb[i]
	}

	return prob
}

// AddProbVal 加入權重參數
func (p *ProbWeight) AddProbVal(v int64) {
	p.probSlice = append(p.probSlice, v)
	p.total += v
}

// GetIndexByProb 按權重取出index
func (p ProbWeight) GetIndexByProb() int {
	if p.total <= 0 || len(p.probSlice) == 0 {
		log.Panic("prob total value is zero")
	}
	var rslt int
	randVal := RandIntTn[int64](p.total)

	for i := 0; i < len(p.probSlice); i++ {
		if randVal >= p.probSlice[i] {
			randVal -= p.probSlice[i]
			continue
		}
		rslt = i
		break
	}

	return rslt
}

type TargetProb[T sliceType] struct {
	targetSlice []T
	ProbWeight
}

// InitTargetProb 初始化權重
func (tp *TargetProb[T]) InitTargetProb(inTarget []T, inProb []int64) {
	if len(inTarget) != len(inProb) {
		log.Panic("target and prob length not equal ")
	}

	tp.targetSlice = append(tp.targetSlice, inTarget...)

	tp.ProbWeight = CreateProb(inProb...)
}

// SetTargetSlice 設定目標參數
func (tp *TargetProb[T]) SetTargetSlice(inTarget []T) {
	tp.targetSlice = make([]T, len(inTarget))
	copy(tp.targetSlice, inTarget)
}

// GetTargetSlice 取得目標參數
func (tp TargetProb[T]) GetTargetSlice() []T {
	return tp.targetSlice
}

// GetOneTargetByProb 依權重取出一對應的目標參數
func (tp TargetProb[T]) GetOneTargetByProb() (int, T) {

	idx := tp.ProbWeight.GetIndexByProb()

	return idx, tp.targetSlice[idx]
}
