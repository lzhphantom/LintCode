package _35

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	result := new(ListNode)
	changeValue(head, result)
	return result
}

func reverseII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		newHead := reverseII(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}
}

var count int

func changeValue(head *ListNode, copy *ListNode) *ListNode {
	var child *ListNode
	if head.Next != nil {
		count++
		child = changeValue(head.Next, copy)
	} else {
		child = copy
	}
	child.Val = head.Val
	if count != 0 {
		count--
		child.Next = &ListNode{}
	}

	return child.Next

}
