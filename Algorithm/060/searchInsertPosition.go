package _60

//给定一个排序数组和一个目标值，如果在数组中找到目标值则返回索引。如果没有，返回到它将会被按顺序插入的位置。
//你可以假设在数组中无重复元素。
func searchInsertPosition(arr []int, target int) int {
	result := len(arr)
	for index, value := range arr {
		if value == target || value > target {
			result = index
			break
		}
	}

	return result
}
