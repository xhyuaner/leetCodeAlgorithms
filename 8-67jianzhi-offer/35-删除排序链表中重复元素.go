package jianzhi

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur := head, head.Next
	for cur != nil {
		if cur.Val == pre.Val {
			cur = cur.Next
		} else {
			pre.Next = cur
			pre = cur
		}
	}
	pre.Next = cur
	return head
}
