package _02

//计算出n阶乘中尾部零的个数
func trailingZeros(n int64) int64 {
	var sum int64 = 1
	for n > 1 {
		sum *= n
		n--
	}
	return countZeros(sum)
}

//计算尾部零个数
func countZeros(n int64) int64 {
	if n == 0 {
		return 1
	}
	var count int64 = 0
	for {
		if n%10 == 0 {
			count++
			n /= 10
		} else {
			break
		}
	}
	return count
}
