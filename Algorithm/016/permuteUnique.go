package _16

import (
	"fmt"
	"reflect"
)

func permuteUnique(nums []int) {
	permAll(nums, 0)
}

func permAll(list []int, k int) {
	if k == len(list)-1 {
		fmt.Println(list)
	} else {
		for i := k; i < len(list); i++ {
			if reflect.DeepEqual(list[i], list[k]) && i != k {
				continue
			}
			swap(&list[i], &list[k])
			permAll(list, k+1)
			swap(&list[i], &list[k])
		}
	}
}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
