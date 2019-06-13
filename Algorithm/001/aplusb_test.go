package _01

import (
	"reflect"
	"testing"
)

func TestAplusb1(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{23, 47, 70},
		{123456789, 987654321, 1111111110},
	}
	for _, test := range tests {
		result := aplusb1(test.a, test.b)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("aplusb1(%d,%d) => %d,want %d", test.a, test.b, result, test.want)
		}
	}
}

func TestAplusb2(t *testing.T) {

	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{23, 47, 70},
		{123456789, 987654321, 1111111110},
	}
	for _, test := range tests {
		result := aplusb2(test.a, test.b)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("aplusb1(%d,%d) => %d,want %d", test.a, test.b, result, test.want)
		}
	}
}
