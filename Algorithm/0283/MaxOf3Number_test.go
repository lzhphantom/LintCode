/*
@Time : 2020/4/5 23:49
@Author : lzhphantom
@File : MaxOf3Number_test
@Software: GoLand
*/
package _283

import (
	"reflect"
	"testing"
)

func TestMaxOfThreeNumbers(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
		num3 int
		want int
	}{
		{1, 2, 3, 3},
		{11, 2, 3, 11},
		{1, 12, 3, 12},
		{111, 2, 3, 111},
	}
	for _, test := range tests {
		result := maxOfThreeNumbers(test.num1, test.num2, test.num3)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("maxOfThreeNumbers(%d,%d,%d)=>%d,want %d", test.num1, test.num2, test.num3, result, test.want)
		}
	}
}
