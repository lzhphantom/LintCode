/*
@Time : 2020/5/29 20:50
@Author : lzhphantom
@File : legalNumberStatistics
@Software: GoLand
*/
package _1656

import (
	"reflect"
	"testing"
)

func TestGetNum(t *testing.T) {
	tests := []struct {
		A    []int
		L    int
		R    int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			3,
			5,
			3,
		},
		{
			[]int{1, 5, 3, 3, 3, 2},
			1,
			2,
			2,
		},
	}

	for _, test := range tests {
		result := getNum(test.A, test.L, test.R)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("test is %d, want is %d", result, test.want)
		}
	}
}
