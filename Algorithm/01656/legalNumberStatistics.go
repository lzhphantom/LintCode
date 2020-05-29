/*
@Time : 2020/5/29 20:50
@Author : lzhphantom
@File : legalNumberStatistics
@Software: GoLand
*/
package _1656

import "sort"

/**
 * @param a: the array a
 * @param L: the integer L
 * @param R: the integer R
 * @return: Return the number of legal integers
 */
func getNum(a []int, L int, R int) int {
	// Write your code here
	sort.Ints(a)
	result := 0
	for _, val := range a {
		if L <= val && val <= R {
			result++
		} else if val > R {
			break
		}
	}
	return result
}
