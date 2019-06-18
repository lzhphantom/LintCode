package _28

import (
	"reflect"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			3,
			true,
		},
	}

	for _, test := range tests {
		result := searchMatrix(test.matrix, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("searchMatrix(%v,%d) => %t,want %t", test.matrix, test.target, result, test.want)
		}
	}
}
