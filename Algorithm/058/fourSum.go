package _58

import (
	"reflect"
	"sort"
)

//Time Limit Exceeded
func fourSum(arr []int, target int) [][]int {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	result := make([][]int, 0)
	startA := 0
	startB := startA + 1
	startC := startB + 1
	startD := startC + 1
out:
	for startA < len(arr) {
		for startD == len(arr) {
			startC++
			startD = startC + 1
		}
		if startD > len(arr) {
			startD = -1
		}
		for startC == len(arr) {
			startB++
			startC = startB + 1
			if startD == -1 {
				startD = startC + 1
				if startD == len(arr) {
					startA++
					startB = startA + 1
					startC = startB + 1
					startD = startC + 1
				}
			}
		}
		if startC > len(arr) {
			startC = -1
		}
		for startB == len(arr) {
			startA++
			startB = startA + 1
			if startC == -1 {
				startC = startB + 1

			}
		}
		if startD == len(arr) {
			break out
		}
		if arr[startA]+arr[startB]+arr[startC]+arr[startD] == target {
			check([]int{arr[startA], arr[startB], arr[startC], arr[startD]}, &result)
		}
		startD++
	}
	return result
}

//for test
func fourSumOther(arr []int, target int) [][]int {
	result := make([][]int, 0)

	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				for l := k + 1; l < len(arr); l++ {
					if arr[i]+arr[j]+arr[k]+arr[l] == target {
						check([]int{arr[i], arr[j], arr[k], arr[l]}, &result)
					}
				}
			}
		}
	}

	return result
}

func check(arr []int, list *[][]int) {
	var count int
	for i := 0; i < len(*list); i++ {
		curArr := append([]int{}, (*list)[i]...)

		if len(curArr) == len(arr) {
			count = 0
			start := 0
			end := len(arr) - 1
			for start != end {
				if curArr[start] == arr[start] {
					count++
				}
				start++
			}
		} else {
			continue
		}

		if count == len(arr) {
			break
		}
	}
	if count != len(arr) {
		isExist := false
		for _, val := range *list {
			if reflect.DeepEqual(val, arr) {
				isExist = true
			}
		}
		if !isExist {
			*list = append(*list, arr)
		}

	}
}
