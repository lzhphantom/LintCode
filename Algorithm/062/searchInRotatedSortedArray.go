package _62

import "sort"

//假设有一个排序的按未知的旋转轴旋转的数组(比如，0 1 2 4 5 6 7 可能成为4 5 6 7 0 1 2)。给定一个目标值进行搜索，
//如果在数组中找到目标值返回数组中的索引位置，否则返回-1。你可以假设数组中不存在重复的元素。

func searchInRotatedSortedArray(arr []int, target int) int {

	result := -1
	if sort.IntsAreSorted(arr) && len(arr) > 0 {
		min := 0
		max := len(arr) - 1
		if target < arr[min] || target > arr[max] {
			return result
		}
		if target == arr[min] {
			return min
		}
		if target == arr[max] {
			return max
		}
		for max-min != 1 {
			middle := (max + min) / 2
			if arr[middle] > target {
				max = middle
			} else if arr[middle] < target {
				min = middle
			} else {
				return middle
			}
		}
	} else {
		for index, value := range arr {
			if value == target {
				result = index
			}
		}
	}

	return result

}
