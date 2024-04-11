package main

import "fmt"

type LinkNode struct {
	value int
	next  *LinkNode
}

func mylist(a *LinkNode, b *LinkNode) int {
	var stackA []int
	var stackB []int
	for a != nil {
		stackA = append(stackA, a.value)
		a = a.next
	}
	for b != nil {
		stackB = append(stackB, b.value)
		b = b.next
	}
	pre := stackA[len(stackA)-1]
	for i, j := len(stackA)-1, len(stackB)-1; i > 0 && j > 0; i, j = i-1, j-1 {
		if stackA[i] != stackB[j] {
			return pre
		}
		pre = stackA[i]
	}
	return pre
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{8, 3, 4, 5}
	a := &LinkNode{}
	headA := a
	b := &LinkNode{}
	headB := b
	for i := 0; i < len(A); i++ {
		a.next = &LinkNode{value: A[i]}
		a = a.next
	}
	for i := 0; i < len(B); i++ {
		b.next = &LinkNode{value: B[i]}
		b = b.next
	}
	fmt.Println(mylist(headA.next, headB.next))

}
