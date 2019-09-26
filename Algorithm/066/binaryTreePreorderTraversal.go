package _66

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	resutlt := make([]int, 0)
	var trees []*TreeNode
	if root == nil {
		return resutlt
	} else {
		trees = append(trees, root)
	}
	for len(trees) > 0 {
		tmp := trees[0]
		resutlt = append(resutlt, tmp.Val)
		num := 0
		if tmp.Left != nil {
			tree := append(trees[:1], tmp.Left)
			trees = append(tree, trees[1:]...)
			num++
		}
		if tmp.Right != nil {
			if num == 0 {
				tree := append(trees[:1], tmp.Right)
				trees = append(tree, trees[1:]...)
			} else {
				tree := append(trees[:2], tmp.Right)
				trees = append(tree, trees[2:]...)
			}
		}
		trees = trees[1:]
	}

	return resutlt
}
