package _53

import (
	"reflect"
	"testing"
)

func TestReverseWordsInAString(t *testing.T) {
	tests := []struct {
		words string
		want  string
	}{
		//{
		//	"the sky is blue",
		//	"blue is sky the",
		//},
		//{
		//	"I'm fine, and you?",
		//	"you? and fine, I'm",
		//},
		{
			"the sky is blue ",
			"blue is sky the",
		},
		{
			" the sky is blue ",
			"blue is sky the",
		},

		{
			" the  sky  is  blue ",
			"blue is sky the",
		},
	}

	for _, test := range tests {
		result := reverseWordsInAString(test.words)

		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("reverseWordsInAString(%s) => %s ,want %s", test.words, result, test.want)
		}
	}
}
