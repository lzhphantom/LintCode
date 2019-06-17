package _16

import "reflect"

var FinalList [][]int

func permuteUnique(nums []int) {
	permAll(nums, 0)
}

func permAll(list []int, k int) {
	if k == len(list)-1 {
		var newlist []int
		if len(FinalList) > 0 {
			notExistFlag := true
			for i := 0; i < len(FinalList); i++ {
				if reflect.DeepEqual(FinalList[i], list) {
					notExistFlag = false
					break
				}
			}
			if notExistFlag {
				newlist = append([]int{}, list...)
				FinalList = append(FinalList, newlist)
			}

		} else {
			newlist = append([]int{}, list...)
			FinalList = append(FinalList, newlist)
		}

	} else {
		for i := k; i < len(list); i++ {
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
