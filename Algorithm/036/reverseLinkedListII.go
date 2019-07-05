package _36

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode, m, n int) *ListNode {
	endListnode := head
	var num int
	// 统计链表长度
	for endListnode != nil {
		num++
		endListnode = endListnode.Next
	}

	if m == n {
		return head
	} else if m > num || n > num || m < 0 || n < 0 {
		return nil
	}
	var max int
	var min int
	if m > n {
		max = m - 1
		min = n - 1
	} else {
		max = n - 1
		min = m - 1
	}
	//存储要交换两个数的地址
	var maxvalue *int
	var minvalue *int
	current := head
	for max > 0 {
		if min == 0 {
			minvalue = &current.Val
		}
		current = current.Next
		max--
		min--
	}
	//交换两个地址
	maxvalue = &current.Val
	*maxvalue, *minvalue = *minvalue, *maxvalue

	return head
}
