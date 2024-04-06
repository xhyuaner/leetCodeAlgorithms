package jianzhi

import "fmt"

func main() {
	list := []int{10, 20, 30, 40, 50, 60, 70, 80}
	dummy := &ListNode{}
	cur := dummy

	for _, value := range list {
		cur.Next = &ListNode{Val: value}
		cur = cur.Next
	}
	PrintList(dummy.Next)
	PrintList(reverseList(dummy.Next))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	p := head
	for p != nil {
		next := p.Next
		p.Next = pre.Next
		pre.Next = p
		p = next
	}
	return pre.Next
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
