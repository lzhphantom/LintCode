package _41

//给定一个整数数组，找到一个具有最大和的子数组，返回其最大和
func maximumSubarray(arr []int) int {
	var result int

	for i := 1; i < len(arr)+1; i++ {
		for j := 0; j <= len(arr)-i; j++ {
			if arr[j] < 0 {
				continue
			}

			now := getSum(arr[j : j+i])

			if now > result {
				result = now
			}
		}
	}

	return result
}

func otherMaximumSubarray(arr []int) int {
	var result int
	var curmax int

	for _, value := range arr {
		if curmax <= 0 {
			curmax = value
		} else {
			curmax += value
		}

		if result < curmax {
			result = curmax
		}
	}

	return result
}

func getSum(arr []int) (result int) {
	for _, value := range arr {
		result += value
	}
	return
}
