/*
@Time : 2020/5/28 22:30
@Author : lzhphantom
@File : sortIntegers_test
@Software: GoLand
*/
package _463

import (
	"fmt"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	tests := []struct {
		Arr *[]int
	}{
		{
			&[]int{3, 2, 1, 4, 5},
		},
		{
			&[]int{1, 1, 2, 1, 1},
		},
	}

	for _, test := range tests {
		sortIntegers(test.Arr)
		fmt.Println(*test.Arr)
	}
}
