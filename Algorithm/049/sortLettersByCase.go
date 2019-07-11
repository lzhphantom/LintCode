package _49

//给定一个只包含字母的字符串，按照先小写字母后大写字母的顺序进行排序。
// *小写字母或者大写字母他们之间不一定要保持在原始字符串中的相对位置。
func sortLettersByCase(chars *string) {
	var theFirstUpper int
	isMark := false
	for index, value := range *chars {
		if !isMark && value >= 'A' && value <= 'Z' {
			theFirstUpper = index
			isMark = true
		}

		if isMark && value >= 'a' && value <= 'z' {
			*chars = (*chars)[:theFirstUpper] + string(rune(value)) + (*chars)[theFirstUpper:index] + (*chars)[index+1:]
		}
	}
}
