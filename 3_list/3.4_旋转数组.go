package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(list)
	rotate(list, 3)
	fmt.Println(list)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(list []int) {
	for left, right := 0, len(list)-1; left < len(list)/2; left++ {
		list[left], list[right-left] = list[right-left], list[left]
	}
}
