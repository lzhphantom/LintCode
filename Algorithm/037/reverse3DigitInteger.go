package _37

func reverse3DigitInteger(number int) int {
	if number < 100 || number > 999 {
		return 0
	}

	return number/100 + (number/10)%10*10 + (number%100)%10*100
}
