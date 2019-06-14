package _05

import "sort"

func kthLargestElement(n int, nums []int) int {
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-n]
}
