package _48

import (
	"reflect"
	"testing"
)

func TestMajorityNumberIII(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			[]int{3, 1, 2, 3, 2, 3, 3, 4, 4, 4},
			3,
			3,
		},
		{
			[]int{1, 1, 2},
			3,
			1,
		},
		{
			[]int{3, 1, 2, 3, 2, 3, 3, 4, 4, 4},
			4,
			3,
		},
	}

	for _, test := range tests {
		result := majorityNumberIII(test.arr, test.k)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("majorityNumber(%v, %d) => %d,want %d", test.arr, test.k, result, test.want)
		}
	}
}
