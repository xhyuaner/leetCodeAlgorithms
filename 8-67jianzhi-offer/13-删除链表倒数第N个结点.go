package jianzhi

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// 让 slowNode 等于 dummy 而不是等于 head，是为了便于后面删除倒数第K个结点
	slowNode := dummy
	for n > 0 {
		head = head.Next
		n--
	}
	for head != nil {
		head = head.Next
		slowNode = slowNode.Next
	}
	// 程序到这里，slowNode 结点是倒数第N个结点的前一个结点
	slowNode.Next = slowNode.Next.Next
	return dummy.Next
}
