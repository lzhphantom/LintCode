package _59

import (
	"reflect"
	"testing"
)

//给一个包含 n 个整数的数组 S, 找到和与给定整数 target 最接近的三元组，返回这三个数的和。
func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			[]int{2, 7, 11, 15},
			3,
			20,
		},
		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
		{
			[]int{-1, 2, 1, -4, 5, 7, -9, 12},
			-1,
			-1,
		},
	}

	for _, test := range tests {
		result := threeSumClosest(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("threeSumClosest(%v, %d) => %d,want %d", test.arr, test.target, result, test.want)
		}
	}
}
