package _33

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		nq   NQueen
		want [][]string
	}{
		{
			NQueen{
				1, 0, make([][]string, 0),
			},
			[][]string{
				{"Q"},
			},
		},
		{
			NQueen{
				4, 0, make([][]string, 0),
			},
			[][]string{
				{".Q..",
					"...Q",
					"Q...",
					"..Q."},
				{"..Q.",
					"Q...",
					"...Q",
					".Q.."},
			},
		},
	}

	for _, test := range tests {
		test.nq.solve()
		result := test.nq.allLayout

		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("%d => %v, want %v", test.nq.size, result, test.want)
		}
	}

}
