package _57

import "sort"

//给出一个有n个整数的数组S，在S中找到三个整数a, b, c，找到所有使得a + b + c = 0的三元组。
//在三元组(a, b, c)，要求a <= b <= c。
//结果不能包含重复的三元组。
func threeSum(arr []int) [][]int {
	result := make([][]int, 0)
	startA := 0
	startB := startA + 1
	startC := startB + 1
	for startA < len(arr) {
		if startC == len(arr) {
			startB++
			startC = startB + 1
		}
		if startB == len(arr) {
			startA++
			startB = startA + 1
		}
		if arr[startA]+arr[startB]+arr[startC] == 0 {
			newArr := []int{arr[startA], arr[startB], arr[startC]}
			sort.Ints(newArr)
			check(newArr, &result)
		}
		startC++
	}

	return result
}

func check(arr []int, target *[][]int) {
	count := 0
	for i := 0; i < len(*target); i++ {
		curArr := append([]int{}, (*target)[i]...)
		if len(curArr) == len(arr) {
			start := 0
			end := len(arr) - 1
			for start != end {
				if curArr[start] == arr[start] {
					count++
				}
				start++
			}
		}
		if count == len(arr) {
			break
		}
	}
	if count != len(arr) {
		*target = append(*target, arr)
	}
}
