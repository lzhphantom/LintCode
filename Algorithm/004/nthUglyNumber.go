package _04

func uthUglyNumber(n int) int {
	var uglyNumberArray = map[int]int{
		2: 0,
		3: 0,
		5: 0,
	}
	var numberArray = []int{1}

	for len(numberArray) < n {
		num2 := numberArray[uglyNumberArray[2]] * 2
		num3 := numberArray[uglyNumberArray[3]] * 3
		num5 := numberArray[uglyNumberArray[5]] * 5
		minNum := min(num2, num3, num5)
		if minNum == num2 {
			uglyNumberArray[2] += 1
		} else if minNum == num3 {
			uglyNumberArray[3] += 1
		} else if minNum == num5 {
			uglyNumberArray[5] += 1
		}
		if minNum == numberArray[len(numberArray)-1] {
			continue
		} else {
			numberArray = append(numberArray, minNum)
		}
	}

	return numberArray[n-1]
}

func min(a, b, c int) int {
	tmp := a
	if b < tmp {
		tmp = b
	}
	if c < tmp {
		tmp = c
	}
	return tmp
}
