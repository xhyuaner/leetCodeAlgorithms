package main

/**
Tag: 回溯

Description:
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

Example:
输入：head = [1,2,3,4]
输出：[2,1,4,3]

Analysis:
时间复杂度：O(n)
空间复杂度：O(n)
*/

func swapPairs(head *ListNode) *ListNode {
	// 终止条件
	if head == nil || head.Next == nil {
		return head
	}
	hdNext := head.Next
	// 设该次递归时需要交换的两个结点是 head 和 hdNext
	// head 连接后面交换完成的子链表
	head.Next = swapPairs(hdNext.Next)
	// hdNext 连接 head，完成交换
	hdNext.Next = head
	// 返回完成交换的链表
	return hdNext
}
