package _47

import "sort"

// 给定一个整型数组，找出主元素，它在数组中的出现次数严格大于数组元素个数的三分之一。
// *数组中只有唯一的主元素
func majorityElementII(arr []int) int {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	var maxNum int
	var maxCount int
	var curNum = arr[0]
	var count int
	for _, value := range arr {
		if value == curNum {
			count++
		} else {
			if count > len(arr)/3 && count > maxCount {
				maxCount = count
				maxNum = curNum
			}
			curNum = value
			count = 1
		}
	}
	if count > len(arr)/3 && count > maxCount {
		maxCount = count
		maxNum = curNum
	}

	return maxNum
}
