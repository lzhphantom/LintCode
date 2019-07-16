package _56

//给一个整数数组，找到两个数使得他们的和等于一个给定的数 target。
//你需要实现的函数twoSum需要返回这两个数的下标, 并且第一个下标小于第二个下标。注意这里下标的范围是 0 到 n-1。
//假设只有一组答案
func twoSum(arr []int, target int) []int {
	result := []int{-1, -1}

	start := 0

	for start < len(arr) {
		if arr[start] < target && result[0] == -1 {
			result[0] = start
			continue
		}
		if result[0] != -1 {
			if arr[start] == target-arr[result[0]] {
				result[1] = start
				break
			}
			if start == len(arr)-1 {
				start = result[0]
				result[0] = -1
			}
		}
		start++
	}

	return result
}
