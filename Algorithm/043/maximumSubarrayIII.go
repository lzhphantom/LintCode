package _43

import "math"

var resultList [][]int

//输入
//查看差别
//[-1,-2,-3,-100,-1,-50]
//2
//输出
//0
//期望答案
//-2
//给定一个整数数组和一个整数 k，找出 k 个不重叠子数组使得它们的和最大。每个子数组的数字在数组中的位置应该是连续的
func maximumSubarrayIII(arr []int, k int) (result int) {
	if k <= 0 || k > len(arr) {
		return 0
	}
	if resultList == nil || k > len(resultList) {
		resultList = make([][]int, k-1)
	}

	if k != 1 {
		for i := 1; i <= len(arr)-k+1; i++ {
			resultList[k-2] = arr[:i]
			curResult := maximumSubarrayIII(arr[i:], k-1)
			if curResult > result {
				result = curResult
			}
			resultList[k-2] = nil
		}
	} else {
		curResult := getMaxarr(arr)
		for _, list := range resultList {
			curResult += getMaxarr(list)
		}
		if curResult > result {
			result = curResult
		}
	}

	return result

}

func getMaxarr(arr []int) int {
	if arr == nil {
		return 0
	}
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
