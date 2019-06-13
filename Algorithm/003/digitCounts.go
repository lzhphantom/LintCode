package _03

func digitCounts(k, n int) int {
	var count int
	for i := n; i >= 0; i-- {
		j := i
		for j >= 10 {
			if j%10 == k {
				count++
			}
			j /= 10
		}
		if j == k {
			count++
		}
	}
	return count
}
