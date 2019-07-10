package _46

import (
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{
			[]int{1, 1, 1, 1, 2, 2, 2},
			1,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1},
			1,
		},
		{
			[]int{1, 1, 1, 2, 2, 2, 2},
			2,
		},
	}

	for _, test := range tests {
		reslut := majorityElement(test.arr)
		if !reflect.DeepEqual(reslut, test.want) {
			t.Errorf("majorityElement(%v) => %d,want %d", test.arr, reslut, test.want)
		}
	}
}
