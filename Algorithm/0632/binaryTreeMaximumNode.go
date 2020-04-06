/*
@Time : 2020/4/6 13:23
@Author : lzhphantom
@File : binaryTreeMaximumNode
@Software: GoLand
*/
package _632

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a *TreeNode, b *TreeNode) *TreeNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val > b.Val {
		return a
	}
	return b
}

func maxNode(tree *TreeNode) *TreeNode {
	if tree == nil {
		return tree
	}
	left := maxNode(tree.Left)
	right := maxNode(tree.Right)
	return Max(tree, Max(left, right))
}

func maxNodeVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return maxNode(node).Val
}
