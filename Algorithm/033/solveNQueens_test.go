package _33

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{
			1,
			[][]string{
				{"Q"},
			},
		},
		{
			4,
			[][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "...Q", ".Q..", "Q..."},
			},
		},
	}

	for _, test := range tests {
		result := solveNQueens(test.n)

		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("solveNQueens(%d) => %v, want %v", test.n, result, test.want)
		}
	}
}
