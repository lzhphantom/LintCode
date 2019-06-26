package _33

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	for i := 0; i < n; i++ {
		result = append(result, make([]string, n))
	}
	helper(result, 0)

	return result
}

func helper(arr [][]string, row int) {
	if row == len(arr) {
		return
	}

	for i := 0; i < len(arr); i++ {
		if isValid(arr, row, i) {
			arr[row][i] = "Q"
			helper(arr, row+1)
			arr[row][i] = "."
		}
	}
}

func isValid(arr [][]string, row, col int) bool {

	for i := 0; i < row; i++ {
		if arr[i][col] == "Q" {
			return false
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if arr[i][j] == "Q" {
				return false
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col + 1; j < len(arr); j++ {
			if arr[i][j] == "Q" {
				return false
			}
		}
	}
	return true
}
