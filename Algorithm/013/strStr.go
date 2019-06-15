package _13

import (
	"reflect"
	"strings"
)

func strStr(source, target string) int {
	return strings.Index(source, target)
}

func strIndex(source, target string) int {
	if len(source) < len(target) {
		return -1
	}
	index := -1
	if len(source) == len(target) {
		if reflect.DeepEqual(source, target) {
			index = 0
		} else {
			return index
		}
	}
	for i := 0; i < len(source); i++ {
		if reflect.DeepEqual(source[i], target[0]) {
			if len(source)-i >= len(target) {
				if reflect.DeepEqual(source[i:i+len(target)], target) {
					index = i
					break
				}
			} else {
				break
			}
		}
	}

	return index
}
