/*
@Time : 2020/4/6 0:10
@Author : lzhphantom
@File : fibonacci
@Software: GoLand
*/
package _366

func fibonacci(n int) int {
	if n <= 2 {
		return n - 1
	}
	first := 0
	second := 1
	result := 0
	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			result = first + second
			second = result
		} else {
			result = first + second
			first = result
		}
	}

	return result
}
