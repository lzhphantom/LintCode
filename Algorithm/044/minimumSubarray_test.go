package _44

import (
	"reflect"
	"testing"
)

func TestMinimumSubarray(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{1, -1, -2, 1},
			-3,
		},
		{
			[]int{1, -1, -2, 1, -4},
			-6,
		},
	}

	for _, test := range tests {
		result := minimumSubarray(test.arr)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("minimumSubarray(%v) => %d,want %d", test.arr, result, test.want)
		}
	}
}
