/*
@Time : 2020/4/6 0:16
@Author : lzhphantom
@File : bibonacci_test
@Software: GoLand
*/
package _366

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{10, 34},
	}

	for _, test := range tests {
		result := fibonacci(test.n)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("fibonacci(%d) => %d,want %d", test.n, result, test.want)
		}
	}
}
