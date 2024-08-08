package jianzhi

/*
*
  - lengthOfLongestSubstring（滑动窗口）
  - @Description:给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
  - @param s
  - @return int
*/
func lengthOfLongestSubstring(s string) int {
	list := [128]int{}
	ans, left := 0, 0
	for right, c := range s {
		list[c]++
		for list[c] > 1 {
			list[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func myfun(s string) int {
	list := [128]int{}
	ans, left := 0, 0
	for right, c := range s {
		list[c]++
		for list[c] > 1 {
			list[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
