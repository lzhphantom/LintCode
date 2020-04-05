package _38

//输入
//[[1,2,3,9],[2,3,9,10],[9,100,109,200]]
//9
//期望答案
//3
func searchIn2DMatrix(arr [][]int, target int) int {
	var count int

	var row int
	var col int
	for row < len(arr) {
		if arr[row][col] == target {
			count++
			col++
		} else if arr[row][col] < target {
			if col == len(arr[row])-1 {
				row++
				col = 0
			} else {
				col++
			}
		} else {
			row++
			col = 0
		}

	}

	return count
}
