package _33

import (
	"fmt"
	"strconv"
)

//var backUp [][]string
//
//func solveNQueens(n int) [][]string {
//	result := make([][]string, 0)
//	for i := 0; i < n; i++ {
//		resultCp := newStringArr(n)
//		checkQueen(&resultCp, i)
//		if checkQueenResutl(resultCp, n) {
//			result = append(result, changeQueenToR(resultCp))
//		}
//	}
//
//	return result
//}
//
//func checkQueen(queens *[][]string, start int) {
//	row := 0
//	(*queens)[row][start] = "Q"
//	Queenlayout(queens, row, start)
//	row++
//	for row < len(*queens) {
//		nextI := canBeLayoutIndex((*queens)[row])
//		if (*queens)[row][nextI] != "" {
//			break
//		}
//		(*queens)[row][nextI] = "Q"
//		Queenlayout(queens, row, nextI)
//		row++
//	}
//
//}
//
//func canBeLayoutIndex(arr []string) int {
//	index := 0
//	for i := 0; i < len(arr); i++ {
//		if arr[i] == "" {
//			index = i
//		}
//	}
//
//	return index
//}
//
//func Queenlayout(queens *[][]string, row int, start int) {
//	for i := 0; i < len((*queens)[row]); i++ {
//		if i != start && (*queens)[row][i] == "" {
//			(*queens)[row][i] = "."
//		}
//	}
//	for i := row + 1; i < len(*queens); i++ {
//		if (*queens)[i][start] == "" {
//			(*queens)[i][start] = "."
//		}
//	}
//
//	queenRow := row + 1
//	queensCol := start + 1
//
//	for queenRow < len(*queens) && queensCol < len(*queens) {
//		if (*queens)[queenRow][queensCol] == "" {
//			(*queens)[queenRow][queensCol] = "."
//		}
//
//		queenRow++
//		queensCol++
//	}
//
//	queenRow = row + 1
//	queensCol = start - 1
//	for queenRow < len(*queens) && queensCol > 0 {
//		if (*queens)[queenRow][queensCol] == "" {
//			(*queens)[queenRow][queensCol] = "."
//		}
//		queenRow++
//		queensCol--
//	}
//}
//
//func changeQueenToR(queens [][]string) []string {
//	newStrings := make([]string, 0)
//	for i := 0; i < len(queens); i++ {
//		newString := strings.Join(queens[i], "")
//		newStrings = append(newStrings, newString)
//	}
//
//	return newStrings
//}
//
//func checkQueenResutl(queens [][]string, n int) bool {
//	queenCounts := 0
//	for i := 0; i < len(queens); i++ {
//		for j := 0; j < len(queens[i]); j++ {
//			if queens[i][j] == "Q" {
//				queenCounts++
//			}
//		}
//	}
//
//	if queenCounts == n {
//		return true
//	}
//	return false
//}
//
//func newStringArr(n int) [][]string {
//	result := make([][]string, 0)
//	for i := 0; i < n; i++ {
//		result = append(result, make([]string, n))
//	}
//	return result
//}

type NQueen struct {
	size      int
	solutions int
	allLayout [][]string
}

func (nq *NQueen) solve() {
	positon := make([]int, nq.size)
	nq.putQueen(positon, 0)
	//fmt.Println("Found", nq.solutions, "solutions.")
}

func (nq *NQueen) putQueen(position []int, target_row int) {
	if target_row == nq.size {
		nq.showFullBoard(position)
		nq.solutions++
	} else {
		for col := 0; col < nq.size; col++ {
			if nq.checkPlace(position, target_row, col) {
				position[target_row] = col
				nq.putQueen(position, target_row+1)
			}
		}
	}
}

func (nq *NQueen) checkPlace(position []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if position[i] == col || position[i]-i == col-row || position[i]+i == col+row {
			return false
		}
	}
	return true
}

func (nq *NQueen) showFullBoard(position []int) {
	newStrings := make([]string, 0)
	for row := 0; row < nq.size; row++ {
		line := ""
		for col := 0; col < nq.size; col++ {
			if position[row] == col {
				line += "Q"
			} else {
				line += "."
			}
		}
		//fmt.Println(line)
		newStrings = append(newStrings, line)
	}
	nq.allLayout = append(nq.allLayout, newStrings)
}

func (nq *NQueen) showShortBoard(position []int) {
	line := ""
	for i := 0; i < nq.size; i++ {
		line += strconv.Itoa(position[i])
	}
	fmt.Println(line)
}
