package _11

import "sort"

type TreeNode struct {
	Val         int
	left, right *TreeNode
}

func searchRange(root *TreeNode, k1, k2 int) []int {
	list := make([]int, 0)
	if root == nil {
		return list
	}
	compareSearchRange(root, &list, k1, k2)
	sort.Ints(list)
	return list
}

func compareSearchRange(root *TreeNode, list *[]int, k1, k2 int) {
	if root.Val >= k1 && root.Val <= k2 {
		*list = append(*list, root.Val)
	}
	if root.left != nil {
		compareSearchRange(root.left, list, k1, k2)
	}
	if root.right != nil {
		compareSearchRange(root.right, list, k1, k2)
	}
}
