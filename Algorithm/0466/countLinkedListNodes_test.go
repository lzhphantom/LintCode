/*
@Time : 2020/4/6 11:14
@Author : lzhphantom
@File : countLinkedListNodes
@Software: GoLand
*/
package _466

import (
	"reflect"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		node *ListNode
		want int
	}{
		{
			&ListNode{
				Val:  1,
				Next: nil,
			},
			1,
		},
		{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			3,
		},
		{
			nil,
			0,
		},
	}
	for _, test := range tests {
		result := countNodes(test.node)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("countNodes(%v)=>%d,want %d", test.node, result, test.want)
		}
	}
}
