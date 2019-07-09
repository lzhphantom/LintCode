package _44

func minimumSubarray(arr []int) int {
	var min int
	var curmin int
	for _, value := range arr {
		if curmin > 0 {
			curmin = value
		} else {
			curmin += value
		}

		if curmin < min {
			min = curmin
		}
	}
	return min
}
