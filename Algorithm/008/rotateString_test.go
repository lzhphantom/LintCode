package _08

import (
	"reflect"
	"testing"
)

func TestRotateString(t *testing.T) {
	tests := []struct {
		str    string
		offset int
		want   string
	}{
		{"abcdefg", 1, "gabcdef"},
		{"abcdefg", 0, "abcdefg"},
		{"abcdefg", 3, "efgabcd"},
		{"abcdefg", 2, "fgabcde"},
	}

	for _, test := range tests {
		rotateString(&test.str, test.offset)
		if !reflect.DeepEqual(test.str, test.want) {
			t.Errorf("rotateString %d => %s,want %s", test.offset, test.str, test.want)
		}
	}
}
