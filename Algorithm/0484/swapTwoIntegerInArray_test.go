/*
@Time : 2020/4/6 10:50
@Author : lzhphantom
@File : swapTwoIntegerInArray
@Software: GoLand
*/
package _484

import (
	"reflect"
	"testing"
)

func TestSwapIntegers(t *testing.T) {
	tests := []struct {
		arr    *[]int
		index1 int
		index2 int
		want   *[]int
	}{
		{&[]int{1, 3, 2}, 1, 2, &[]int{1, 2, 3}},
		{&[]int{1, 2, 3, 4}, 1, 2, &[]int{1, 3, 2, 4}},
		{&[]int{1, 2, 3, 4}, 2, 3, &[]int{1, 2, 4, 3}},
		{&[]int{1, 2, 2, 2}, 0, 3, &[]int{2, 2, 2, 1}},
	}

	for _, test := range tests {
		arr := make([]int, len(*test.arr))
		copy(arr, *test.arr)
		swapIntegers(test.arr, test.index1, test.index2)
		if !reflect.DeepEqual(test.arr, test.want) {
			t.Errorf("swapIntegers(%v,%d,%d)=>%v,want %v", arr, test.index1, test.index2, test.arr, test.want)
		}
	}
}
