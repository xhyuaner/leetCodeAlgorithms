package KMP

/**
Tag: KMP算法

Description:
给你两个字符串 haystack 和 needle ，
请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。

Example:
输入：haystack = "sadbutsad", needle = "sad"
输出：0
输入：haystack = "leetcode", needle = "leeto"
输出：-1

Analysis:
时间复杂度：O(m + n)
空间复杂度：O(1)
*/

func strStr(haystack string, needle string) int {
	m := len(needle)
	if m == 0 {
		return 0
	}
	j := 0
	next := make([]int, m)
	getNext(next, needle) // 获取next数组
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到 j 的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// 获取无减一或者右移的next数组（前缀表）abcabcc	0001230
func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
