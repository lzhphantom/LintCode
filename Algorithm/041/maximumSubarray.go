package _41

//给定一个整数数组，找到一个具有最大和的子数组，返回其最大和
func maximumSubarray(arr []int) int {
	var result int

	for _, value := range arr {
		if value >= 0 {
			result += value
		}
	}

	return result
}
