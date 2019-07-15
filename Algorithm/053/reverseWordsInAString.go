package _53

import "strings"

//单词的构成：无空格字母构成一个单词，有些单词末尾会带有标点符号
//输入字符串是否包括前导或者尾随空格？可以包括，但是反转后的字符不能包括
//如何处理两个单词间的多个空格？在反转字符串中间空格减少到只含一个
func reverseWordsInAString(words string) string {
	wordList := strings.Split(words, " ")

	for i := 0; i < len(wordList); i++ {
		if wordList[i] == "" {
			wordList = append(wordList[:i], wordList[i+1:]...)
		}
	}

	reverseString(wordList)

	newString := strings.Join(wordList, " ")

	return newString

}

func reverseString(arr []string) {
	start := 0
	end := len(arr) - 1

	for end-start >= 1 {
		swap(&arr[start], &arr[end])
		start++
		end--
	}
}

func swap(a, b *string) {
	var temp string
	temp = *a
	*a = *b
	*b = temp
}
