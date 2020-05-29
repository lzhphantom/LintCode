/*
@Time : 2020/5/29 21:24
@Author : lzhphantom
@File : pascalsTriangle
@Software: GoLand
*/
package _1355

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{
			1,
			[][]int{{1}},
		},
		{
			2,
			[][]int{{1}, {1, 1}},
		},
		{
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}
	for _, test := range tests {
		result := generate(test.numRows)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("test is %v, want is %v", result, test.want)
		}
	}
}
