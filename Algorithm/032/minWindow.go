package _32

import "strings"

func minWindow(source, target string) string {
	if len(source) < len(target) {
		return ""
	}

	start := make([]int, 0)

	for index, value := range source {
		if strings.ContainsRune(target, value) {
			start = append(start, index)
		}
	}

	var result string

	if len(start) < len(target) {
		return result
	}

	for j := len(target); j <= len(source); j++ {
		for i := 0; i < len(start); i++ {
			if start[i]+j > len(source) {
				break
			}
			result = source[start[i] : start[i]+j]

			rightNum := 0
			for _, value := range target {
				if strings.ContainsRune(result, value) {
					rightNum++
				}
			}
			if rightNum == len(target) {
				goto endfor
			}
		}
	}

endfor:

	return result

}
