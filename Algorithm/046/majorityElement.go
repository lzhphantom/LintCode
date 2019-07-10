package _46

import "sort"

// 给定一个整型数组，找出主元素，它在数组中的出现次数严格大于数组元素个数的二分之一。
// *假设数组非空，且数组中总是存在主元素。
func majorityElement(arr []int) int {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	var curNum = arr[0]
	var count int
	for _, value := range arr {
		if value == curNum {
			count++
		} else {
			if count > len(arr)/2 {
				break
			}
			curNum = value
			count = 1
		}
	}

	return curNum
}
