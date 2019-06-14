package _05

import "sort"

func insertSort(nums []int) []int {
	var tmp []int
	for _, num := range nums {
		if len(tmp) == 0 {
			tmp = append(tmp, num)
			continue
		}
		for index, t := range tmp {
			if num > t {
				if index+1 == len(tmp) {
					tmp = append(tmp, num)
					break
				}
			} else {
				tmp = insert(tmp, num, index)
			}
		}
	}
	return tmp
}

func insert(nums []int, n, index int) []int {
	s := append([]int{n}, nums[index:]...)
	nums = append(nums[:index], s...)
	return nums
}

func kthLargestElement(n int, nums []int) int {
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-n]
}
