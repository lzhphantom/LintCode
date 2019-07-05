package _38

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		arr    [][]int
		target int
		want   int
	}{
		{
			[][]int{
				{1, 3, 4, 7},
				{2, 4, 7, 8},
				{3, 5, 9, 10},
			},
			3,
			2,
		},
	}

	for _, test := range tests {
		result := searchIn2DMatrix(test.arr, test.target)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("Ssearchin2DMatrix(%v, %d) => %d,want %d", test.arr, test.target, result, test.want)
		}
	}
}
