/*
@Time : 2020/4/6 10:50
@Author : lzhphantom
@File : swapTwoIntegerInArray
@Software: GoLand
*/
package _484

/**
 * @param A: An integer array
 * @param index1: the first index
 * @param index2: the second index
 * @return: nothing
 */
func swapIntegers(A *[]int, index1 int, index2 int) {
	(*A)[index1], (*A)[index2] = (*A)[index2], (*A)[index1]
}
