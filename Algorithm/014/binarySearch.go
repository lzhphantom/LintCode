package _14

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	index := -1
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		index = end
	}

	for end-start != 1 {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			index = mid
			end--
		}
	}

	return index
}
