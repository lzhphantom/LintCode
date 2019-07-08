package _42

//给定一个整数数组，找出两个 不重叠 子数组使得它们的和最大
//每个子数组的数字在数组中的位置应该是连续的
//返回最大的和
func maximumSubarrayII(arr []int) int {
	theFirstArrSum := 0
	col := 1
	start := 0
	theSecondStart := 0
	for col < len(arr) {
		if start+col > len(arr) {
			col++
			start = 0
			continue
		}
		if arr[start] < 0 {
			start++
			continue
		}
		now := getSum(arr[start : start+col])
		if theFirstArrSum < now {
			theFirstArrSum = now
			theSecondStart = start + col
		}
		start++
		if start == len(arr) {
			col++
			start = 0
		}

	}
	col = 1
	start = theSecondStart
	theSecondArrSum := 0
	for col <= len(arr)-theSecondStart {
		if start+col > len(arr) {
			col++
			start = theSecondStart
			continue
		}
		if arr[start] < 0 {
			start++
			continue
		}
		now := getSum(arr[start : start+col])
		if theSecondArrSum < now {
			theSecondArrSum = now
		}
		start++
		if start == len(arr) {
			col++
			start = theSecondStart
		}
	}

	return theFirstArrSum + theSecondArrSum
}

func getSum(arr []int) (result int) {
	for _, value := range arr {
		result += value
	}
	return
}
