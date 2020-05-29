/*
@Time : 2020/5/29 20:07
@Author : lzhphantom
@File : jewelsAndStones
@Software: GoLand
*/
package _1038

import (
	"reflect"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{
			"aA",
			"aAAbbbb",
			3,
		},
		{
			"z",
			"ZZ",
			0,
		},
	}
	for _, test := range tests {
		result := numJewelsInStones(test.J, test.S)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("test is %v, want is %v", result, test.want)
		}
	}
}
