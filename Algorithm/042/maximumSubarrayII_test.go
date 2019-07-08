package _42

import (
	"reflect"
	"testing"
)

func TestMaximumSubarrayII(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{1, 3, -1, 2, -1, 2},
			7,
		},
		{
			[]int{5, 4},
			9,
		},
	}

	for _, test := range tests {
		result := maximumSubarrayII(test.arr)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("maximumSubarrayII(%v) => %d, want %d", test.arr, result, test.want)
		}
	}
}
