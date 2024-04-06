package main

/**
 * main
 *  @Description: 求一个字符串中同时包含‘r'和’e'但不包含‘d'字符的连续子串的个数
 */

import (
	"fmt"
)

func countSubstrings(s string) int {
	// 初始化记录位置的变量，-1表示还未找到
	lastR := -1
	lastE := -1
	lastD := -1
	count := 0

	for i, char := range s {
		// 更新最近一次出现r、e、d的位置
		switch char {
		case 'r':
			lastR = i
		case 'e':
			lastE = i
		case 'd':
			lastD = i
		}

		// 如果'r'和'e'都已出现至少一次，并且它们的位置都在最近一次'd'出现的位置之后
		if lastR != -1 && lastE != -1 && lastR > lastD && lastE > lastD {
			// 计算满足条件的新子串数，取lastR和lastE中较小的那个与lastD之间的距离
			count += min(lastR, lastE) - lastD
		}
	}

	return count
}

// min 返回两个整数中较小的一个
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var s string
	fmt.Println("输入字符串:")
	fmt.Scan(&s)

	result := countSubstrings(s)
	fmt.Println("满足条件的连续子串数量:", result)
}
