package _45

import (
	"reflect"
	"testing"
)

func TestMaximumSubarrayDifference(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{1, 2, -3, 1},
			6,
		},
		{
			[]int{0, -1},
			1,
		},
	}

	for _, test := range tests {
		result := maximumSubarrayDifference(test.arr)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("maximumSubarrayDifference(%v) => %d,want %d", test.arr, result, test.want)
		}
	}
}
