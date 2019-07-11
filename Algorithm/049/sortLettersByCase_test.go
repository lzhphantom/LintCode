package _49

import (
	"reflect"
	"testing"
)

func TestSortLettersByCase(t *testing.T) {
	tests := []struct {
		chars string
		want  string
	}{
		{
			"abAcD",
			"abcAD",
		},
		{
			"ABC",
			"ABC",
		},
		{
			"abAcBadD",
			"abdacABD",
		},
	}

	for _, test := range tests {

		sortLettersByCase(&test.chars)
		if !reflect.DeepEqual(test.chars, test.want) {
			t.Errorf("sortLettersByCase => %s,want %s", test.chars, test.want)
		}
	}
}
