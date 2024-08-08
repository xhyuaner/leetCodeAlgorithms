package jianzhi

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	dummy := &ListNode{}
	h := dummy
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			h.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			h.Next = pHead2
			pHead2 = pHead2.Next
		}
		h = h.Next
	}
	if pHead1 != nil {
		h.Next = pHead1
	}
	if pHead2 != nil {
		h.Next = pHead2
	}
	return dummy.Next
}
