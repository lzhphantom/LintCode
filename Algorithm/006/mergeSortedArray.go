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

func mergeSortedArray2(arr1, arr2 []int) []int {
	if arr1[0] >= arr2[len(arr2)-1] {
		return append(arr2, arr1...)
	} else if arr2[0] >= arr1[len(arr1)-1] {
		return append(arr1, arr2...)
	}
	countArr1 := 0
	countArr2 := 0
	var tmp = make([]int, 0)
	for countArr1 != len(arr1) && countArr2 != len(arr2) {
		if arr1[countArr1] < arr2[countArr2] {
			tmp = append(tmp, arr1[countArr1])
			countArr1++
		} else if arr1[countArr1] == arr2[countArr2] {
			tmp = append(tmp, arr1[countArr1], arr2[countArr2])
			countArr1++
			countArr2++
		} else {
			tmp = append(tmp, arr2[countArr2])
			countArr2++
		}
	}
	if countArr1 == len(arr1) {
		tmp = append(tmp, arr2[countArr2:]...)
	}
	if countArr2 == len(arr2) {
		tmp = append(tmp, arr1[countArr1:]...)
	}
	return tmp
}
