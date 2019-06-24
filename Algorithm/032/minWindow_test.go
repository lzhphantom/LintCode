package _32

import (
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		source string
		target string
		want   string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"",
			"",
			"",
		},
		{
			"CBA",
			"ABC",
			"CBA",
		},
		{
			"ABC",
			"N",
			"",
		},
		{
			"ABC",
			"NC",
			"",
		},
	}

	for _, test := range tests {
		result := minWindow(test.source, test.target)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("minWindow(%s, %s) => %s, want %s", test.source, test.target, result, test.want)
		}
	}
}
