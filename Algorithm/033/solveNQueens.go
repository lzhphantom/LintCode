package _33

import (
	"fmt"
	"strconv"
)

type NQueen struct {
	size      int
	solutions int
	allLayout [][]string
}

func (nq *NQueen) solve() {
	position := make([]int, nq.size)
	nq.putQueen(position, 0)
	//fmt.Println("Found", nq.solutions, "solutions.")
}

// 通过检查所有N个可能的情况，尝试在targetRow上放置一个女王。
// 如果找到一个有效的位置，该函数会调用自己试图在下一行放置一个女王，
// 直到所有N个女王被放置在NxN板上。
func (nq *NQueen) putQueen(position []int, targetRow int) {
	if targetRow == nq.size {
		nq.showFullBoard(position)
		nq.solutions++
	} else {
		for col := 0; col < nq.size; col++ {
			if nq.checkPlace(position, targetRow, col) {
				position[targetRow] = col
				nq.putQueen(position, targetRow+1)
			}
		}
	}
}

// 检查给定位置是否受到任何先前放置的皇后的攻击（检查列和对角线位置）
func (nq *NQueen) checkPlace(position []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if position[i] == col || /*同一列是否有皇后*/
			position[i]-i == col-row || /*正斜角上是否有皇后*/
			position[i]+i == col+row { /*反斜角上是否有皇后*/
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
