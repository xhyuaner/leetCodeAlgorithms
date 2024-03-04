package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}
	// 计算需要反转的结点数量
	n := right - left + 1
	if n == 1 { // left=right，只反转一个结点，相当于没变，直接返回
		return head
	}
	// 在头结点前添加一个哑结点
	preNode := &ListNode{Next: head}
	// 让头结点指针指向这个哑结点，便于后续返回反转后链表的头结点
	head = preNode
	// 找到反转前左边界结点的前一个结点，便于后续采用头插法进行反转
	for ; left >= 2; left-- {
		preNode = preNode.Next
	}
	// 使用一个右指针来指向反转后边界结点，便于后续将反转区间后面的结点连接到链表尾部
	rightNode := preNode.Next
	// 采用头插法反转区间链表
	leftNode := preNode.Next
	nextNode := leftNode.Next
	preNode.Next = nil
	for ; n > 0; n-- {
		leftNode.Next = preNode.Next
		preNode.Next = leftNode
		leftNode = nextNode
		if nextNode != nil {
			nextNode = nextNode.Next
		}
	}
	// 将反转区间后面的结点连接到链表尾部
	rightNode.Next = leftNode
	return head.Next
}
