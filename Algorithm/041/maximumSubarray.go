package _41

import "math"

//给定一个整数数组，找到一个具有最大和的子数组，返回其最大和
func maximumSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	n := len(arr)
	global := make([]int, 2)
	local := make([]int, 2)
	global[0] = arr[0]
	local[0] = arr[0]

	for i := 1; i < n; i++ {
		local[i%2] = Max(arr[i], local[(i-1)%2]+arr[i])
		global[i%2] = Max(local[i%2], global[(i-1)%2])
	}
	return global[(n-1)%2]
}

func otherMaximumSubarray(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	max := math.MinInt32
	sum := 0
	for _, val := range arr {
		sum += val
		max = Max(max, sum)
		sum = Max(sum, 0)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSum(arr []int) (result int) {
	for _, value := range arr {
		result += value
	}
	return
}
