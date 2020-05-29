/*
@Time : 2020/5/29 21:08
@Author : lzhphantom
@File : numberOf1Bits
@Software: GoLand
*/
package _1332

import "fmt"

/**
 * @param n: an unsigned integer
 * @return: the number of ’1' bits
 */
func hammingWeight(n int) int {
	// write your code here
	bits := fmt.Sprintf("%b", n)
	result := 0
	for _, val := range bits {
		if val == '1' {
			result++
		}
	}
	return result
}
