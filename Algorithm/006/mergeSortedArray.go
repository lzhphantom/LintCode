package _06

import "sort"

func mergeSortedArray(arr1, arr2 []int) []int {
	if arr1[0] >= arr2[len(arr2)-1] {
		return append(arr2, arr1...)
	} else if arr2[0] >= arr1[len(arr1)-1] {
		return append(arr1, arr2...)
	}
	tmp := append(arr1, arr2...)
	sort.Ints(tmp)
	return tmp
}
