package _11

import "sort"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
	if root.Left != nil {
		compareSearchRange(root.Left, list, k1, k2)
	}
	if root.Right != nil {
		compareSearchRange(root.Right, list, k1, k2)
	}
}

/**
 * @param root: param root: The root of the binary search tree
 * @param k1: An integer
 * @param k2: An integer
 * @return: return: Return all keys that k1<=key<=k2 in ascending order
 */
func searchRange1(root *TreeNode, k1 int, k2 int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val >= k1 && root.Val <= k2 {
			return []int{root.Val}
		}
		return []int{}
	}
	result := make([]int, 0)
	if root.Left != nil && root.Val > k1 {
		result = append(result, searchRange(root.Left, k1, k2)...)
	}
	if root.Val >= k1 && root.Val <= k2 {
		result = append(result, root.Val)
	}
	if root.Right != nil && root.Val < k2 {
		result = append(result, searchRange(root.Right, k1, k2)...)
	}
	return result
}
