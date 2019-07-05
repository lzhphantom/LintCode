package _37

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{123, 321},
		{234, 432},
		{100, 1},
		{900, 9},
	}

	for _, test := range tests {
		result := reverse3DigitInteger(test.num)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("reverse3DigitInteger(%d) => %d,want %d", test.num, result, test.want)
		}
	}
}
