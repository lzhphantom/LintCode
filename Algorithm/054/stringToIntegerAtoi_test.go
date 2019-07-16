package _54

import (
	"math"
	"reflect"
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			"123",
			123,
		},
		{
			"12.3",
			12,
		},
		{
			"12312312312321312312312321321",
			math.MaxInt32,
		},
		{
			"2147483648",
			math.MaxInt32,
		},
		{
			"+2147483648",
			math.MaxInt32,
		},
		{
			"-2147483648",
			math.MinInt32,
		},
		{
			"-2147483647",
			-2147483647,
		},
		{
			"-123",
			-123,
		},
	}

	for _, test := range tests {
		result := atoi(test.str)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("atoi(%s) => %d, want %d", test.str, result, test.want)
		}
	}
}
