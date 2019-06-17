package _20

var S = make([]int, 0)

func dicesSum(n int) map[int]float64 {
	dices := []int{1, 2, 3, 4, 5, 6}
	sumList(dices, n, 0)
	result := make(map[int]int)
	for i := 0; i < len(S); i++ {
		if _, ok := result[S[i]]; ok {
			result[S[i]] += 1
		} else {
			result[S[i]] = 1
		}
	}
	valueWithProbability := make(map[int]float64)
	for key, value := range result {
		valueWithProbability[key] = float64(value) / float64(len(S))
	}
	return valueWithProbability
}

func sumList(sum []int, n int, total int) {
	for i := 0; i < len(sum); i++ {
		total += sum[0]
		if n > 1 {
			sumList(sum, n-1, total)
		} else {
			S = append(S, total)
		}
	}
}
