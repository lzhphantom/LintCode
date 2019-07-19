package _61

//给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
//如果目标值不在数组中，则返回[-1, -1]
func searchForARange(arr []int, target int) []int {
	result := []int{-1, 1}
	record := 0
	for index, value := range arr {
		if record >= len(result) {
			break
		}
		if value == target {
			result[record] = index
			record++
		}
	}
	return result
}
