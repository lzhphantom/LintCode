package _36

type ListNode struct {
	Val  int
	Next *ListNode
}

//查看差别
//3760->2881->7595->3904->5069->4421->8560->8879->8488->5040->5792->56->1007->2270->3403->6062->null
//2
//7
//输出
//3760->8560->7595->3904->5069->4421->2881->8879->8488->5040->5792->56->1007->2270->3403->6062->null
//期望答案
//3760->8560->4421->5069->3904->7595->2881->8879->8488->5040->5792->56->1007->2270->3403->6062->null
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
