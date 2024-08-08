package jianzhi

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 指定 当前结点 和 它的前一个结点
	pre, cur := head, head.Next
	for cur != nil {
		// 如果当前结点和前一个结点相同，就让当前结点一直++，直到遇到与pre结点不同的结点
		if cur.Val == pre.Val {
			cur = cur.Next
		} else { // 将前面的重复结点一次性删除，让pre=cur，但是cur此时不用++,因为下次循环会++
			pre.Next = cur
			pre = cur
		}
	}
	// 将pre结点的后一个结点置空，应对[11233]场景
	pre.Next = cur
	return head
}
