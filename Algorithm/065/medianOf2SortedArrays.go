package _65

func medianOf2SortedArrays(arr1, arr2 []int) interface{} {
	var result interface{}
	if len(arr1) == 0 {
		return medofArray(arr2)
	}
	if len(arr2) == 0 {
		return medofArray(arr1)
	}
	if (len(arr1)+len(arr2))%2 == 0 {

	} else {
		theNum := (len(arr1) + len(arr2) - 1) / 2
		startarr1 := 0
		startarr2 := 0
		for theNum > 0 {
			if arr1[startarr1] < arr2[startarr2] {
				result = float32(arr1[startarr1])
			} else {
				result = float32(arr2[startarr2])
			}
			theNum--
		}

	}

	return result
}

func medofArray(arr []int) interface{} {
	if len(arr) == 0 {
		return -1
	}
	return (arr[len(arr)/2] + arr[(len(arr)-1)/2]) / 2
}
