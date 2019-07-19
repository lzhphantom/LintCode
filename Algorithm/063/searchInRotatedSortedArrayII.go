package _63

import "sort"

//跟进“搜索旋转排序数组”，假如有重复元素又将如何
//写出一个函数判断给定的目标值是否出现在数组中
func searchInRotatedSortedArrayII(arr []int, target int) bool {
	result := false
	if sort.IntsAreSorted(arr) && len(arr) > 0 {
		min := 0
		max := len(arr) - 1
		if target < arr[min] || target > arr[max] {
			return result
		}
		if target == arr[min] || target == arr[max] {
			return true
		}
		for max-min != 1 {
			middle := (max + min) / 2
			if arr[middle] > target {
				max = middle
			} else if arr[middle] < target {
				min = middle
			} else {
				return true
			}
		}
	} else {
		for _, value := range arr {
			if value == target {
				result = true
				break
			}
		}
	}
	return result
}
