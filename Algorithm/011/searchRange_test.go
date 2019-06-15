package _11

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		k1, k2 int
		want   []int
	}{
		{
			&TreeNode{
				3,
				&TreeNode{
					4,
					nil, nil,
				},
				&TreeNode{
					8,
					nil, nil,
				},
			},
			3, 10,
			[]int{3, 4, 8},
		},
		{
			&TreeNode{
				1,
				&TreeNode{
					5,
					&TreeNode{
						Val: 9,
					},
					&TreeNode{
						Val: 20,
					},
				},
				&TreeNode{
					13,
					&TreeNode{
						Val: 10,
					},
					&TreeNode{
						Val: 31,
					},
				},
			},
			7, 14,
			[]int{9, 10, 13},
		},
	}

	for _, test := range tests {
		list := searchRange(test.root, test.k1, test.k2)
		if !reflect.DeepEqual(list, test.want) {
			t.Errorf("searchRange(%v,%d,%d) => %v,want %v", test.root, test.k1, test.k2, list, test.want)
		}
	}
}

func TestSearchRange1(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		k1, k2 int
		want   []int
	}{
		{
			&TreeNode{
				3,
				&TreeNode{
					4,
					nil, nil,
				},
				&TreeNode{
					8,
					nil, nil,
				},
			},
			3, 10,
			[]int{3, 4, 8},
		},
		{
			&TreeNode{
				1,
				&TreeNode{
					5,
					&TreeNode{
						Val: 9,
					},
					&TreeNode{
						Val: 20,
					},
				},
				&TreeNode{
					13,
					&TreeNode{
						Val: 10,
					},
					&TreeNode{
						Val: 31,
					},
				},
			},
			7, 14,
			[]int{9, 10, 13},
		},
	}

	for _, test := range tests {
		list := searchRange1(test.root, test.k1, test.k2)
		if !reflect.DeepEqual(list, test.want) {
			t.Errorf("searchRange1(%v,%d,%d) => %v,want %v", test.root, test.k1, test.k2, list, test.want)
		}
	}
}
