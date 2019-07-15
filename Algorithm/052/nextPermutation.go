package _52

func nextPermutation(arr []int) {

	if arr == nil || len(arr) <= 1 {
		return
	}

	// 先从后往前找到第一个不是依次增长的数
	// 记录下位置i
	i := len(arr) - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	// 如果i存在，从i开始往后找，找到下一个最接近i上的数且比i上的数大，然后两个调换位置
	if i >= 0 {
		j := i + 1
		for j < len(arr) && arr[j] > arr[i] {
			j++
		}
		j--
		swap(&arr[i], &arr[j])
	}
	// i之后的数，倒序
	l := i + 1
	reverseInt(arr[l:])
}

func reverseInt(arr []int) {
	start := 0
	end := len(arr) - 1

	for end-start >= 1 {
		swap(&arr[start], &arr[end])
		start++
		end--
	}
}

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
