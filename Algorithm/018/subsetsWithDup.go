package _18

import "reflect"

func subsetsWithDup(nums []int) [][]int {
	list := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		if i == len(nums) {
			list = append(list, []int{})
		} else if i == 0 {
			list = append(list, nums)
		} else {
			for j := 0; j < len(nums); j++ {
				newlist := append([]int{}, nums...)
				count := i
				for count != 0 {
					newlist = delList(newlist, j)
					count--
				}

				list = listAdd(list, newlist)
			}

		}
	}
	return list
}

func listAdd(list [][]int, nums []int) [][]int {
	notExistFlag := true
	for i := 0; i < len(list); i++ {
		if reflect.DeepEqual(list[i], nums) {
			notExistFlag = false
			break
		}
	}
	if notExistFlag {
		list = append(list, nums)
	}
	return list
}

func delList(nums []int, index int) []int {
	if index >= len(nums) {
		index %= len(nums)
	}
	list := append([]int{}, nums[:index]...)
	nums = append(list, nums[index+1:]...)
	return nums
}
