package _15

import "fmt"

func permute(nums []int) {
	permAll(nums, 0)
}

func permAll(list []int, k int) {
	if k == len(list)-1 {
		fmt.Println(list)
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
