package jianzhi

/**
 * rotateRight
 *  @Description: 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置
 *  @param head
 *  @param k
 *  @return *ListNode
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 统计链表的长度
	len := 1
	iter := head
	for iter.Next != nil {
		len++
		iter = iter.Next
	}
	// 计算出还需要几个结点
	need := len - k%len // k%len：计算出倒数第几个结点作为头结点
	if need == len {
		return head
	}
	// 将链表成环
	iter.Next = head
	// 找到需要断开的位置
	for need > 0 {
		iter = iter.Next
		need--
	}
	// 找到旋转后的结果头结点
	ret := iter.Next
	// 将结果头结点前面的连接断开
	iter.Next = nil
	return ret
}
