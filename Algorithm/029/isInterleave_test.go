package _29

import (
	"reflect"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1, s2, s3 string
		want       bool
	}{
		{"abcd", "abcd", "aabbccdd", true},
		{"", "", "1", false},
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
	}

	for _, test := range tests {
		result := isInterleave(test.s1, test.s2, test.s3)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("isInterleave(%s, %s, %s) => %t, want %t", test.s1, test.s2, test.s3, result, test.want)
		}
	}
}
