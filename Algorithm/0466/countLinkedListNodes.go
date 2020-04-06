/*
@Time : 2020/4/6 11:14
@Author : lzhphantom
@File : countLinkedListNodes
@Software: GoLand
*/
package _466

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	tmp := l
	output := ""
	for tmp != nil {
		output += fmt.Sprintf("%d->", tmp.Val)
		tmp = tmp.Next
	}
	return output + "null"
}

/**
 * @param head: the first node of linked list.
 * @return: An integer
 */
func countNodes(head *ListNode) int {
	// write your code here
	if head == nil {
		return 0
	}
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
