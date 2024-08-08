package sliding_window

/*
*
Tag: 滑动窗口 + 哈希表

Description:
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

Example:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

Analysis:
时间复杂度：O(n+m+Σ)，其中 n 为字符串 s 的长度，m 为字符串 p 的长度，其中Σ 为所有可能的字符数。
我们需要 O(m) 来统计字符串 p 中每种字母的数量；需要 O(m) 来初始化滑动窗口；需要 O(Σ) 来初始化 differ；需要 O(n−m) 来滑动窗口
并判断窗口内每种字母的数量是否与字符串 p 中每种字母的数量相同，每次判断需要 O(1) 。因为 s 和 p 仅包含小写字母，所以 Σ=26。
空间复杂度：O(Σ)。用于存储滑动窗口和字符串 p 中每种字母数量的差。
*/
func findAnagrams1(s string, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for j, c := range s[:sLen-pLen] {
		sCount[c-'a']--
		sCount[s[j+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, j+1)
		}
	}
	return
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	count := [26]int{}
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		if count[ch-'a'] == 1 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[ch-'a'] == 0 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[ch-'a']--

		if count[s[i+pLen]-'a'] == -1 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[s[i+pLen]-'a'] == 0 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[s[i+pLen]-'a']++

		if differ == 0 {
			ans = append(ans, i+1)
		}
	}
	return
}
