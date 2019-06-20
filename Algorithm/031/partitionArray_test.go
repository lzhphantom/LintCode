package _31

import (
	"reflect"
	"testing"
)

func TestPartitionArray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{}, 9, 0},
		{[]int{3, 2, 2, 1}, 2, 1},
		{[]int{3, 9, 6, 4, 2, 10, 8, 7, 13, 17, 19, 20}, 7, 4},
	}

	for _, test := range tests {
		result := partitionArray(test.nums, test.k)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("partitionArray(%v, %d) => %d, want %d", test.nums, test.k, result, test.want)
		}
	}
}
