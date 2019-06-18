package _28

//写出一个高效的算法来搜索 m × n矩阵中的值。
//
//这个矩阵具有以下特性：
//
//每行中的整数从左到右是排序的。
//每行的第一个数大于上一行的最后一个整数。
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		list := matrix[i]
		if target <= list[len(list)-1] {
			for _, value := range list {
				if value == target {
					return true
				}
			}
			break
		}
	}
	return false
}
