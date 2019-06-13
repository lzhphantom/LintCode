package _03

import (
	"reflect"
	"testing"
)

func TestDigitCounts(t *testing.T) {
	tests := []struct {
		k, n int
		want int
	}{
		{1, 12, 5},
		{0, 12, 2},
		{1, 100, 21},
		{2, 100, 20},
		{3, 100, 20},
		{4, 100, 20},
		{5, 100, 20},
		{6, 100, 20},
		{7, 100, 20},
		{8, 100, 20},
		{9, 100, 20},
		{0, 100, 12},
	}

	for _, test := range tests {
		result := digitCounts(test.k, test.n)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("digitCounts(%d, %d) => %d,want %d", test.k, test.n, result, test.want)
		}
	}
}
