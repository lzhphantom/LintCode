package _09

import "strconv"

func fizzBuzz(n int) []string {
	list := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			list = append(list, "fizz")
		} else if i%3 != 0 && i%5 == 0 {
			list = append(list, "buzz")
		} else if i%3 == 0 && i%5 == 0 {
			list = append(list, "fizz buzz")
		} else {
			list = append(list, strconv.Itoa(i))
		}
	}
	return list
}
