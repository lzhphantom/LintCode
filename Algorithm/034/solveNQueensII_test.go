package _34

import (
	"reflect"
	"testing"
)

func TestNQueensII(t *testing.T) {
	tests := []struct {
		nq   NQueens
		want int
	}{
		{
			NQueens{
				1,
				0,
			},
			1,
		},
		{
			NQueens{
				4,
				0,
			},
			2,
		},
		{
			NQueens{
				7,
				0,
			},
			40,
		},
	}

	for _, test := range tests {
		test.nq.solve()
		if !reflect.DeepEqual(test.nq.solutions, test.want) {
			t.Errorf("%d NQueens have %d ways,want %d ways", test.nq.size, test.nq.solutions, test.want)
		}
	}
}
