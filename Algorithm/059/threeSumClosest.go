package _59

import (
	"math"
	"sort"
)

func threeSumClosest(arr []int, target int) int {
	if len(arr) == 3 {
		return arr[0] + arr[1] + arr[2]
	}

	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	startA := 0
	startB := startA + 1
	startC := startB + 1

	min := math.MaxInt32
	var sum int
	for startA < len(arr)-3 {
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
				if startC == len(arr) {
					break
				}
			}
		}

		if abs(arr[startA]+arr[startB]+arr[startC]-target) < min {
			min = abs(arr[startA] + arr[startB] + arr[startC] - target)
			sum = arr[startA] + arr[startB] + arr[startC]
		}
		startC++
	}
	return sum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
