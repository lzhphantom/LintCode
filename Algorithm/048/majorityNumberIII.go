package _48

import "sort"

// 给定一个整型数组，找到主元素，它在数组中的出现次数严格大于数组元素个数的1/k。
// *数组中只有唯一的主元素
func majorityNumberIII(arr []int, k int) int {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	var maxNumber int
	var maxCount int
	curNumber := arr[0]
	var curCount int

	for _, value := range arr {
		if value == curNumber {
			curCount++
		} else {
			if curCount > len(arr)/k && curCount > maxCount {
				maxCount = curCount
				maxNumber = curNumber
			}
			curNumber = value
			curCount = 1
		}
	}
	if curCount > len(arr)/3 && curCount > maxCount {
		maxCount = curCount
		maxNumber = curNumber
	}

	return maxNumber
}
