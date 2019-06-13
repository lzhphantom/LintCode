package _02

import (
	"reflect"
	"testing"
)

func TestTrailingZeros(t *testing.T) {
	tests := []struct {
		n, count int64
	}{
		{0, 0},
		{11, 2},
		{12, 2},
		{17, 3},
	}

	for _, test := range tests {
		result := trailingZeros(test.n)

		if !reflect.DeepEqual(result, test.count) {
			t.Errorf("trailingZeros(%d) => %d,want %d", test.n, result, test.count)
		}
	}
}
