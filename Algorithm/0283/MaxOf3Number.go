/*
@Time : 2020/4/5 23:45
@Author : lzhphantom
@File : MaxOf3Number
@Software: GoLand
*/
package _283

/**
 * @param num1: An integer
 * @param num2: An integer
 * @param num3: An integer
 * @return: an interger
 */
func maxOfThreeNumbers(num1 int, num2 int, num3 int) int {
	if num1 > num2 {
		if num1 > num3 {
			return num1
		} else {
			return num3
		}
	} else {
		if num2 > num3 {
			return num2
		} else {
			return num3
		}
	}
}
