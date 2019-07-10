package _45

import "math"

// 给定一个整数数组，找出两个不重叠的子数组A和B，使两个子数组和的差的绝对值|SUM(A) - SUM(B)|最大。
func maximumSubarrayDifference(arr []int) int {
	var result int
	for i := 1; i < len(arr); i++ {
		for start := 0; start < len(arr)-i; start++ {
			curmax := getMax(arr[start : start+i])
			curmin := getMin(arr[start+i:])

			if result < curmax-curmin {
				result = curmax - curmin
			}
		}
	}
	return result
}

func getMax(arr []int) int {
	var max = math.MinInt32
	var curMax int
	for _, value := range arr {
		if curMax <= 0 {
			curMax = value
		} else {
			curMax += value
		}

		if max < curMax {
			max = curMax
		}
	}

	return max
}

func getMin(arr []int) int {
	var min = math.MaxInt32
	var curmin int
	for _, value := range arr {
		if curmin > 0 {
			curmin = value
		} else {
			curmin += value
		}

		if min > curmin {
			min = curmin
		}
	}
	return min
}
