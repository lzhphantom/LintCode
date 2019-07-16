package _54

import (
	"math"
	"strings"
)

//实现atoi这个函数，将一个字符串转换为整数。如果没有合法的整数，返回0。如果整数超出了32位整数的范围，返回INT_MAX(2147483647)如果是正整数，或者INT_MIN(-2147483648)如果是负整数。
func atoi(str string) int {

	if ((strings.Compare(str, "2147483647") > 0 || len(str) > 10) && (str[0] >= '0' && str[0] <= '9')) || (str[0] == '+' && (strings.Compare(str[1:], "2147483647") > 0 || len(str) > 11)) {
		return math.MaxInt32
	} else if str[0] == '-' && len(str) > 11 {
		return math.MinInt32
	} else if str[0] == '-' && len(str) == 11 {
		minNum := "2147483648"
		start := 0
		end := len(minNum)
		for start != end {
			if start != end-1 && str[start+1] > minNum[start] {
				return math.MinInt32
			}
			if start != end-1 && str[start+1] < minNum[start] {
				break
			}
			if start == end-1 && str[start+1] >= minNum[start] {
				return math.MinInt32
			}
			start++
		}
	}

	var sum int
	str0 := str
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}

	for _, value := range []byte(str) {
		if value == '.' {
			break
		}
		num := value - '0'
		sum = sum*10 + int(num)
	}

	if str0[0] == '-' {
		sum = -sum
	}
	return sum
}
