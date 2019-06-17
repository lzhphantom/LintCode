package _17

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {

	fmt.Println(subsets([]int{1, 2, 3, 4}))
	//[[1 2 3 4] [2 3 4] [1 3 4] [1 2 4] [1 2 3] [3 4] [1 4] [1 2] [2 3] [4] [1] [2] [2] []]

}
