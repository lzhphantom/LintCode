/*
@Time : 2020/5/29 21:08
@Author : lzhphantom
@File : numberOf1Bits
@Software: GoLand
*/
package _1332

import (
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{11, 3},
		{7, 3},
	}
	for _, test := range tests {
		result := hammingWeight(test.n)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("test is %v,want is %v", result, test.want)
		}
	}
}
