package _57

import "sort"

//给出一个有n个整数的数组S，在S中找到三个整数a, b, c，找到所有使得a + b + c = 0的三元组。
//在三元组(a, b, c)，要求a <= b <= c。
//结果不能包含重复的三元组。
func threeSum(arr []int) [][]int {
	result := make([][]int, 0)

	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	startA := 0
	startB := startA + 1
	startC := startB + 1
bad:
	for startA < len(arr) {
		for startC == len(arr) {
			startB++
			startC = startB + 1
		}
		if startC > len(arr) {
			startC = -1
		}
		for startB == len(arr) {
			startA++
			startB = startA + 1
			if startC == -1 {
				startC = startB + 1
				if startC >= len(arr) {
					break bad
				}
			}
		}
		if arr[startA]+arr[startB]+arr[startC] == 0 {
			newArr := []int{arr[startA], arr[startB], arr[startC]}
			check(newArr, &result)
		}
		startC++
	}

	return result
}

//for test
func threeSumOther(arr []int) [][]int {
	result := make([][]int, 0)
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					check([]int{arr[i], arr[j], arr[k]}, &result)
				}
			}
		}
	}

	return result
}

func check(arr []int, target *[][]int) {
	var count int
	for i := 0; i < len(*target); i++ {
		curArr := append([]int{}, (*target)[i]...)
		if len(curArr) == len(arr) {
			count = 0
			start := 0
			end := len(arr)
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
		*target = append(*target, arr)
	}
}
