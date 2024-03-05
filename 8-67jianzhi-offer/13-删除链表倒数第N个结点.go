package jianzhi

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 如果链表为空，或者只有一个结点，直接返回删除倒数第一个结点的结果nil
	if head == nil || head.Next == nil {
		return nil
	}
	// 使用preNode结点，便于删除第N个结点
	preNode := &ListNode{Next: head}
	slowNode := head
	fastNode := head
	for ; n > 0; n-- {
		fastNode = fastNode.Next
	}
	if fastNode == nil {
		return preNode.Next.Next
	}
	for fastNode != nil {
		preNode = slowNode
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	preNode.Next = preNode.Next.Next
	return head
}
