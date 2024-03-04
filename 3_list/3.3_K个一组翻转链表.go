package main

/**
Tag: 模拟

Description:
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

Example:
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

Analysis:
时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{
		Next: head,
	}
	// pre 标记上一组的最后一个结点
	pre := hair
	// tail 不停向后探索的结点
	tail := pre
	// nex 记录下一组的第一个结点
	var nex *ListNode
	for head != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex = tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	p := head
	var nex *ListNode
	for pre != tail {
		nex = p.Next
		p.Next = pre
		pre = p
		p = nex
	}
	return tail, head
}

/*func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	emptyHeadNode := &ListNode{}
	preHead := emptyHeadNode
	// 标记每一组的最后一个结点
	endNode := preHead
	// 标记每一组开头结点的下一个结点
	var headNext *ListNode
	// 定义一个检测 K 个结点是否都不为空的结点
	var isNilNode *ListNode
	var isOk bool
	for head != nil {
		isOk = true
		isNilNode = head
		// 判断剩余结点能否构成一组
		for i := 0; i < k; i++ {
			if isNilNode != nil {
				isNilNode = isNilNode.Next
			} else {
				isOk = false
			}
		}
		// 如果剩余结点能构成一组，就执行翻转
		if isOk {
			endNode = head
			for i := 0; i < k; i++ {
				headNext = head.Next
				head.Next = preHead.Next
				preHead.Next = head
				head = headNext
			}
			preHead = endNode
		} else { // 否则就结束循环
			preHead.Next = head
			break
		}
	}
	return emptyHeadNode.Next
}*/
