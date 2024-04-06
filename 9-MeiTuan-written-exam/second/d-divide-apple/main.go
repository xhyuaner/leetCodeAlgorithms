package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	// 每个小孩至少分到一个苹果
	remainingApples := m - n

	// 计算小明最多可以分到多少苹果
	maxApplesForXiaoming := maxApples(k, n, remainingApples) + 1 // 加1是因为每个小孩初始至少分到一个苹果

	fmt.Println(maxApplesForXiaoming)
}

func maxApples(k, n, remainingApples int) int {
	left := min(k-1, n-k) // 小明左边和右边小孩的最小数量
	right := n - 1 - left // 小明另一侧的小孩数量
	rounds := 0           // 完整的左右分配轮次

	// 尝试分配剩余的苹果
	for remainingApples > 0 {
		// 需要分配的苹果数量：左边、右边加上小明自己
		required := 2*left + right + 1
		if remainingApples >= required {
			remainingApples -= required
			rounds++
		} else {
			break
		}
	}

	// 小明可以额外获得的苹果数是分配轮次的数量
	return rounds
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
