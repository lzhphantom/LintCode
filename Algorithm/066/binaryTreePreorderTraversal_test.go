package _66

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		//{
		//	&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}},
		//	[]int{1, 2, 3},
		//},
		//{
		//	&TreeNode{1, nil, &TreeNode{2, &TreeNode{Val: 3}, nil}},
		//	[]int{1, 2, 3},
		//},
		{
			&TreeNode{1, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 5}}, &TreeNode{Val: 3}},
			[]int{1, 2, 4, 5, 3},
		},
	}

	for _, test := range tests {
		result := preorderTraversal(test.root)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("preorderTraversal(%v) => %v,want %v", test.root, result, test.want)
		}
	}
}
