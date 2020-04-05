package _17

//有问题
//输入
//查看差别
//[0,0]
//输出
//[[0,0],[0],[0],[]]
//期望答案
//[[],[0],[0],[0,0]]
func subsets(nums []int) [][]int {
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

				list = append(list, newlist)
			}

		}
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
