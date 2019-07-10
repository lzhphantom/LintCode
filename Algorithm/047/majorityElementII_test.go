package _47

import (
	"reflect"
	"testing"
)

func TestMajorityElementII(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{99, 2, 99, 2, 99, 3, 3},
			99,
		},
		{
			[]int{1, 2, 1, 2, 1, 3, 3},
			1,
		},
	}

	for _, test := range tests {
		reslut := majorityElementII(test.arr)
		if !reflect.DeepEqual(reslut, test.want) {
			t.Errorf("majorityElement(%v) => %d,want %d", test.arr, reslut, test.want)
		}
	}
}
