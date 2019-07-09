package _43

import "math"

var resultList [][]int
var final int

//给定一个整数数组和一个整数 k，找出 k 个不重叠子数组使得它们的和最大。每个子数组的数字在数组中的位置应该是连续的
func maximumSubarrayIII(arr []int, k int) {
	if k <= 0 || k > len(arr) {
		return
	}
	if resultList == nil {
		resultList = make([][]int, k-1)
	}

	if k != 1 {
		for i := 1; i < len(arr)-k+1; i++ {
			resultList[k-2] = arr[:i]
			maximumSubarrayIII(arr[i:], k-1)
		}
	} else {
		result := getMaxarr(arr)
		for _, list := range resultList {
			result += getMaxarr(list)
		}
		if result > final {
			final = result
		}
	}

}

func getMaxarr(arr []int) int {
	sum := math.MinInt32
	var cursum int
	for _, value := range arr {
		if cursum <= 0 {
			cursum = value
		} else {
			cursum += value
		}
		if cursum > sum {
			sum = cursum
		}
	}
	return sum
}
