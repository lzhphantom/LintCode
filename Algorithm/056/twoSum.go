package _56

//给一个整数数组，找到两个数使得他们的和等于一个给定的数 target。
//你需要实现的函数twoSum需要返回这两个数的下标, 并且第一个下标小于第二个下标。注意这里下标的范围是 0 到 n-1。
//假设只有一组答案
func twoSum(arr []int, target int) []int {
	result := []int{-1, -1}
	first := 0
	second := 1
	for first < len(arr) && second < len(arr) {
		if arr[first]+arr[second] == target {
			result[0] = first
			result[1] = second
			break
		} else {
			second++
			if second >= len(arr) {
				first++
				second = first + 1
			}
		}
	}

	return result
}
