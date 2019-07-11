package _48

import "sort"

func majorityNumberIII(arr []int, k int) int {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	var maxNumber int
	var maxCount int
	curNumber := arr[0]
	var curcount int

	for _, value := range arr {
		if value == curNumber {
			curcount++
		} else {
			if curcount > len(arr)/k && curcount > maxCount {
				maxCount = curcount
				maxNumber = curNumber
			}
			curNumber = value
			curcount = 1
		}
	}

	return maxNumber
}
