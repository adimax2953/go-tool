package gotool

// 氣泡排序Sli(小到大)
func BubbleSort(list []int32) []int32 {
	var temp int32
	var i int
	var j int
	for i = 0; i < len(list)-1; i++ {
		for j = len(list) - 1; j > i; j-- {
			if list[j-1] > list[j] {
				temp = list[j-1]
				list[j-1] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}

//region 快速排序
func division(list []int32, left int32, right int32) int32 {

	// 以最左邊的數(left)為基準
	var base int32 = list[left]
	for left < right {
		// 從序列右端開始, 向左遍歷, 直到找到小於base的數
		for left < right && list[right] >= base {
			right--
		}
		// 找到了比base小的元素, 將這個元素放到最左邊的位置
		list[left] = list[right]
		// 從序列左端開始, 向右遍歷, 直到找到大於base的數
		for left < right && list[left] <= base {
			left++
		}
		// 找到了比base大的元素, 將這個元素放到最右邊的位置
		list[right] = list[left]

	}
	// 最後將base放到left位置。此時, left位置的左側數值應該都比left小
	// 而left位置的右側數值應該都比left大。
	list[left] = base //此時left == right
	return left
}

func QuickSort(list []int32, left int32, right int32) {
	// 左下標一定小於右下標, 否則就越界了
	if left < right {
		//對數組進行分割, 取出下次分割的基準標號
		var base int32 = division(list, left, right)
		//對“基準標號“左側的一組數值進行遞歸的切割, 以至於將這些數值完整的排序
		QuickSort(list, left, base-1)
		//對“基準標號“右側的一組數值進行遞歸的切割, 以至於將這些數值完整的排序
		QuickSort(list, base+1, right)
	}

}
