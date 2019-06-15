package _09

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "fizz"}},
		{6, []string{"1", "2", "fizz", "4", "buzz", "fizz"}},
		{15, []string{"1", "2", "fizz",
			"4", "buzz", "fizz",
			"7", "8", "fizz",
			"buzz", "11", "fizz",
			"13", "14", "fizz buzz"}},
		{20, []string{"1", "2", "fizz",
			"4", "buzz", "fizz",
			"7", "8", "fizz",
			"buzz", "11", "fizz",
			"13", "14", "fizz buzz", "16", "17", "fizz", "19", "buzz"}},
	}

	for _, test := range tests {
		list := fizzBuzz(test.n)
		if !reflect.DeepEqual(list, test.want) {
			t.Errorf("fizzBuzz(%d) => %v,want %v", test.n, list, test.want)
		}
	}
}
