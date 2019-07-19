package _60

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			[]int{1, 3, 5, 6},
			0,
			0,
		},
	}

	for _, test := range tests {
		result := searchInsertPosition(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("searchInsertPosition(%v, %d) => %d,want %d", test.arr, test.target, result, test.want)
		}
	}
}
