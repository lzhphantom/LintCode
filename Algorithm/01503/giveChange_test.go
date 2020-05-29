/*
@Time : 2020/5/29 12:48
@Author : lzhphantom
@File : giveChange_test
@Software: GoLand
*/
package _1503

import (
	"reflect"
	"testing"
)

func TestGiveChange(t *testing.T) {
	tests := []struct {
		amount int
		want   change
	}{
		{
			1004,
			change{
				Four:   1,
				OneSix: 1,
			},
		}, {
			1014,
			change{
				One:  2,
				Four: 2,
			},
		},
	}
	for _, test := range tests {
		result := giveChange(test.amount)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("test is %v,want is %v", result, test.want)
		}
	}
}
