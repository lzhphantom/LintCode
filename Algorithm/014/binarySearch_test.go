package _14

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			4,
			3,
		},
		{
			[]int{1, 2, 3, 3, 4, 5, 10},
			3,
			2,
		},
		{
			[]int{1, 2, 3, 3, 4, 5, 10},
			6,
			-1,
		},
		{
			[]int{1, 2, 3, 3, 4, 5, 10},
			1,
			0,
		},
		{
			[]int{1, 2, 3, 3, 4, 5, 10, 10},
			10,
			6,
		},
	}

	for _, test := range tests {
		result := binarySearch(test.nums, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("binarySearch(%v, %d) => %d,want %d", test.nums, test.target, result, test.want)
		}
	}
}
