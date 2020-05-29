/*
@Time : 2020/5/29 20:07
@Author : lzhphantom
@File : jewelsAndStones
@Software: GoLand
*/
package _1038

/**
 * @param J: the types of stones that are jewels
 * @param S: representing the stones you have
 * @return: how many of the stones you have are also jewels
 */
func numJewelsInStones(J string, S string) int {
	// Write your code here
	result := 0
	for _, sVal := range S {
		for _, jVal := range J {
			if jVal == sVal {
				result++
				break
			}
		}
	}
	return result
}
