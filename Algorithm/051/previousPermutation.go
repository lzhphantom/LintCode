package _51

//给定一个整数数组来表示排列，找出其上一个排列。
//样例
//给出排列[1,3,2,3]，其上一个排列是[1,2,3,3]
//给出排列[1,2,3,4]，其上一个排列是[4,3,2,1]
func perviousPermutation(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
				reverseInt(arr[i+1:])
				return
			}
		}
	}
	reverseInt(arr)
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
