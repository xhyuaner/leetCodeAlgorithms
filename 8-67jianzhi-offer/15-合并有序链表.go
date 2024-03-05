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
			dummy.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			dummy.Next = pHead2
			pHead2 = pHead2.Next
		}
		dummy = dummy.Next
	}
	if pHead1 != nil {
		dummy.Next = pHead1
	}
	if pHead2 != nil {
		dummy.Next = pHead2
	}
	return h.Next
}
