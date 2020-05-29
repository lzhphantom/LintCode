/*
@Time : 2020/5/28 22:27
@Author : lzhphantom
@File : sortIntegers
@Software: GoLand
*/
package _463

/**
 * @param A: an integer array
 * @return: nothing
 */
func sortIntegers(A *[]int) {
	// write your code here
	for i := 0; i < len(*A); i++ {
		for j := 0; j < len(*A)-i-1; j++ {
			if (*A)[j] > (*A)[j+1] {
				tmp := (*A)[j]
				(*A)[j] = (*A)[j+1]
				(*A)[j+1] = tmp
			}
		}
	}
}
