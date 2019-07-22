package _64

//合并两个排序的整数数组A和B变成一个新的数组。
//你可以假设A具有足够的空间（A数组的大小大于或等于m+n）去添加B中的元素。
func mergeSortedArray(arr1 *[]int, m int, arr2 []int, n int) {
	arr2Start := 0
	arr1Start := 0
	for arr1Start < m && arr2Start < n {
		if arr2[arr2Start] < (*arr1)[arr1Start] {
			newArr := append([]int{}, (*arr1)[:arr1Start]...)
			newArr = append(newArr, arr2[arr2Start])
			*arr1 = append(newArr, (*arr1)[arr1Start:]...)
			arr2Start++
		} else {
			arr1Start++
		}
	}
	if arr2Start != n {
		*arr1 = append(*arr1, arr2[arr2Start:]...)
	}
}
