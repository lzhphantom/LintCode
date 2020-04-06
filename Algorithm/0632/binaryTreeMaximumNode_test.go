/*
@Time : 2020/4/6 13:23
@Author : lzhphantom
@File : binaryTreeMaximumNode
@Software: GoLand
*/
package _632

import (
	"reflect"
	"testing"
)

func TestMaxNode(t *testing.T) {
	tests := []struct {
		tree *TreeNode
		want *TreeNode
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -5,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   -4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   -5,
						Left:  nil,
						Right: nil,
					},
				},
			},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   -4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -5,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, test := range tests {
		result := maxNode(test.tree)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("maxNode(%v)=>%v,want%v", test.tree, result, test.want)
		}
	}
}

func TestMaxNodeVal(t *testing.T) {
	tests := []struct {
		tree *TreeNode
		want int
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -5,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   -4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   -5,
						Left:  nil,
						Right: nil,
					},
				},
			},
			3,
		},
	}
	for _, test := range tests {
		result := maxNodeVal(test.tree)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("maxNode(%v)=>%v,want%v", test.tree, result, test.want)
		}
	}
}
