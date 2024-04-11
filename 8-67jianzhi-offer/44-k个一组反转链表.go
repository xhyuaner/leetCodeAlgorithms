package jianzhi

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	tail := hair
	for head != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nxt := tail.Next
		head, tail = reverse1(head, tail)
		pre.Next = head
		tail.Next = nxt
		pre = tail
		head = nxt
	}
	return hair.Next
}

func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	cur := head
	for prev != tail {
		nex := cur.Next
		cur.Next = prev
		prev = cur
		cur = nex
	}
	return tail, head
}
