package _61

import (
	"reflect"
	"testing"
)

func TestSearchForARange(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   []int
	}{
		{
			[]int{},
			9,
			[]int{-1, 1},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
	}

	for _, test := range tests {
		result := searchForARange(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("searchForARange(%v, %d) => %d,want %d", test.arr, test.target, result, test.want)
		}
	}
}
