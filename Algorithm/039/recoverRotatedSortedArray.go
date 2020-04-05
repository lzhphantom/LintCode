package _39

import "sort"

//给定一个旋转排序数组，再恢复其排序
func recoverRotatedSortedArray(arr []int) []int {
	var number int
	var index int
	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] >= number {
			number = arr[i]
		}
		if arr[i] < number {
			index = i
			break
		}
	}

	result := append(arr[index:], arr[:index]...)

	return result
}

func recoverRotatedSortedArrayII(arr *[]int) {
	sort.Ints(*arr)
}
