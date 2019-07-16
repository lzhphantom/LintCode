package _55

import (
	"reflect"
	"testing"
)

func TestCompareStrings(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{
			"ABC",
			"ABC",
			true,
		},
		{
			"ABCD",
			"ABC",
			true,
		},
		{
			"ABCD",
			"AABBCC",
			true,
		},
		{
			"ABCD",
			"AABBCCE",
			false,
		},
	}

	for _, test := range tests {
		result := compareStrings(test.a, test.b)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("atoi(%s,%s) => %v, want %v", test.a, test.b, result, test.want)
		}
	}
}
