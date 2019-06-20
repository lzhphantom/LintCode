package _31

//给出一个整数数组 nums 和一个整数 k。划分数组（即移动数组 nums 中的元素），使得：
//
//所有小于k的元素移到左边
//所有大于等于k的元素移到右边
//返回数组划分的位置，即数组中第一个位置 i，满足 nums[i] 大于等于 k。
func partitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	lgCount := 0
	eqCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= k {
			lgCount++
			if nums[i] == k {
				eqCount++
			}
		}

	}
	return lgCount - eqCount
}
