/*
@Time : 2020/5/29 21:24
@Author : lzhphantom
@File : pascalsTriangle
@Software: GoLand
*/
package _1355

/**
 * @param numRows: num of rows
 * @return: generate Pascal's triangle
 */
func generate(numRows int) [][]int {
	// write your code here
	result := make([][]int, 0)
	if numRows >= 1 {
		result = append(result, []int{1})
	}
	if numRows >= 2 {
		result = append(result, []int{1, 1})
	}

	if numRows > 2 {
		for i := 3; i <= numRows; i++ {
			arr := make([]int, i)
			arr[0] = 1
			arr[i-1] = 1
			for j := 1; j < i-1; {
				for k := 0; k < len(result[i-2])-1; k++ {
					arr[j] = result[i-2][k] + result[i-2][k+1]
					j++
				}
			}
			result = append(result, arr)
		}
	}

	return result
}
