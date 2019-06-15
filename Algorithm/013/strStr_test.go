package _13

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		source, target string
		want           int
	}{
		{"source", "target", -1},
		{"abcdefg", "bcd", 1},
	}

	for _, test := range tests {
		result := strStr(test.source, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("strStr(%s, %s) => %d,want %d", test.source, test.target, result, test.want)
		}
	}
}

func TestStrIndex(t *testing.T) {
	tests := []struct {
		source, target string
		want           int
	}{
		{"source", "target", -1},
		{"abcdefg", "bcd", 1},
	}

	for _, test := range tests {
		result := strIndex(test.source, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("strStr(%s, %s) => %d,want %d", test.source, test.target, result, test.want)
		}
	}
}
