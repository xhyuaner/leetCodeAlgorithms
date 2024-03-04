package backtrack

/**
Tag: 分治 + 回溯

Description:
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

Example:
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]

Analysis:
时间复杂度：O(n*logk)
空间复杂度：O(logk)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsBackTrack(lists, 0, len(lists))
}

// 递归二分合并多个有序链表
func mergeKListsBackTrack(lists []*ListNode, i, j int) *ListNode {
	m := j - i
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[i]
	}
	left := mergeKListsBackTrack(lists, i, i+m/2)
	right := mergeKListsBackTrack(lists, i+m/2, j)
	return mergeTwoLists(left, right)
}

// 合并两个有序链表
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	}
	if list2 != nil {
		prev.Next = list2
	}
	return preHead.Next
}
