package _41

import (
	"reflect"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{-2, 2, -3, 4, -1, 2, 1, -5, 3},
			6,
		},
		{
			[]int{1, 2, 3, 4},
			10,
		},
		{
			[]int{-1},
			-1,
		},
	}

	for _, test := range tests {
		result := maximumSubarray(test.arr)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("maximumSubarray(%v) => %d, want %d", test.arr, result, test.want)
		}
	}
}

func TestOtherMaximumSubarray(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{-2, 2, -3, 4, -1, 2, 1, -5, 3},
			6,
		},
		{
			[]int{1, 2, 3, 4},
			10,
		},
	}

	for _, test := range tests {
		result := otherMaximumSubarray(test.arr)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("maximumSubarray(%v) => %d, want %d", test.arr, result, test.want)
		}
	}
}
