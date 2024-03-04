package main

/**
Tag：链表 + 双指针

给你一个单向链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。、

时间复杂度：O(L)，其中 L 是链表的长度
空间复杂度：O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 在头结点之前添加一个哑结点
	dummy := &ListNode{0, head}
	// first指针指向头结点，second指针指向哑结点
	first, second := head, dummy
	// 先让first结点往前遍历 n 次，此时first和second指针正好相隔 n 个结点
	for i := 0; i < n; i++ {
		first = first.Next
	}
	// 再让first和second结点同时往前遍历，直到first结点到达链表尾部null
	// 此时second结点正好指向倒数第n+1个结点，只需删除它后面那个结点即可
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	// 删除倒数第 n 个结点
	second.Next = second.Next.Next
	// 返回头结点
	return dummy.Next
}
