package gotool

import (
	"reflect"
	"runtime"
)

// SearchSliInt32 - 檢查silce中有沒有此元素(int32)
func SearchSliInt32(slice []int32, elem int32) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

//  SearchSliInt64 - 檢查silce中有沒有此元素(int64)
func SearchSliInt64(slice []int64, elem int64) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// SearchSliInt - 檢查silce中有沒有此元素(int)
func SearchSliInt(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// SearchSliFlt - 檢查silce中有沒有此元素(float64)
func SearchSliFlt(slice []float64, elem float64) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// SearchSliStr - 檢查silce中有沒有此元素(str)
func SearchSliStr(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// RemoveSliInt - 移除slice中的某個元素(int)
func RemoveSliInt(slice []int32, elem int32) []int32 {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...) // 只刪除一個該元素, （唯一的元素）
			// return RemoveEleSlice(slice, elem)// 遞歸刪除全部該元素
			break
		}
	}
	return slice
}

// RemoveSliStr - 移除slice中的某個元素(str)
func RemoveSliStr(slice []string, elem string) []string {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...) // 只刪除一個該元素, （唯一的元素）
			// return RemoveEleSlice(slice, elem)// 遞歸刪除全部該元素
			break
		}
	}
	return slice
}

// int32 slice 加總
func Int32Sum(int32arr []int32) int32 {
	var sum int32 = 0
	for _, i := range int32arr {
		sum += int32(i)
	}
	return sum
}

func Int32SumParallel(numbers []int32) int32 {
	nNum := len(numbers)
	var total int32 = 0
	if nNum < 100000 {
		total = Int32Sum(numbers)
	} else {
		nCPU := runtime.NumCPU()

		ch := make(chan int32)
		for i := 0; i < nCPU; i++ {
			from := i * nNum / nCPU
			to := (i + 1) * nNum / nCPU
			go func() { ch <- Int32Sum(numbers[from:to]) }()
		}
		for i := 0; i < nCPU; i++ {
			total += <-ch
		}
	}

	return total
}

func SliceMaxFloat64(l []float64) (max float64) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return
}

func SliceMinFloat64(l []float64) (min float64) {
	min = l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return
}

func Removeduplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

// 去除重複資料 壓測效能優化
func RemoveduplicateMap(DataArr []int32) []int32 {
	var resultArr []int32
	resultmap := make(map[int32]bool, 3)
	for _, i := range DataArr { //大量或者string可以改用i++減少每次賦值的操作
		if _, ok := resultmap[i]; ok {
			continue
		} else {
			resultmap[i] = true
			resultArr = append(resultArr, i)
		}
	}
	// LogTool.LogDebug("單骰結果:%v", SingleDicArr)
	return resultArr
}

// IntArrEq - 兩個Int arr比較是否相等
func IntArrEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
