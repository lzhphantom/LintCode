package _34

type NQueens struct {
	size      int
	solutions int
}

func (nq *NQueens) solve() {
	position := make([]int, nq.size)
	nq.putQueens(position, 0)
}

func (nq *NQueens) putQueens(position []int, targetRow int) {
	if targetRow == nq.size {
		nq.solutions++
	} else {
		for col := 0; col < nq.size; col++ {
			if checkAround(position, targetRow, col) {
				position[targetRow] = col
				nq.putQueens(position, targetRow+1)
			}
		}
	}
}

func checkAround(position []int, targetRow, col int) bool {
	for row := 0; row < targetRow; row++ {
		if position[row] == col || position[row]-row == col-targetRow || position[row]+row == col+targetRow {
			return false
		}
	}
	return true
}
