package _50

//给定一个整数数组A。
//定义B[i] = A[0] * ... * A[i-1] * A[i+1] * ... * A[n-1]， 计算B的时候请不要使用除法。请输出B。

func productOfArrayExcludeItself(arr []int) []int {
	result := make([]int, len(arr))

	for i := 0; i < len(result); i++ {
		sum := 1
		for index, value := range arr {
			if index == i {
				continue
			}
			sum *= value
		}
		result[i] = sum
	}

	return result
}
