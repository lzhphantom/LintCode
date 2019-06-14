package _04

import (
	"reflect"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{9, 10},
		{10, 12},
	}

	for _, test := range tests {
		result := uthUglyNumber(test.n)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("uthUglyNumber(%d) => %d,want %d", test.n, result, test.want)
		}
	}
}
