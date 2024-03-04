package sliding_window

/**
Tag: 多起点滑动窗口 + 哈希表

Description:
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
 s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。

Example:
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]

Analysis:
时间复杂度：O(ls * n),其中 ls 是输入 s 的长度，n 是 words 中每个单词的长度。
			需要做 n 次滑动窗口，每次需要遍历一次 s
空间复杂度：O(m * n), m 是 words 中单词的个数
*/
func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	// 循环遍历字符串 s n次，因为有 n 个不同的起点
	for i := 0; i < n && i+m*n <= ls; i++ {
		// 初始化哈希表 differ 表示窗口中单词频次和 words 中单词频次之差
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		// 滑动窗口
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				tempWord := s[start+(m-1)*n : start+m*n] // 窗口新加入的单词
				differ[tempWord]++
				if differ[tempWord] == 0 {
					delete(differ, tempWord)
				}
				tempWord = s[start-n : start] // 窗口去除旧的单词
				differ[tempWord]--
				if differ[tempWord] == 0 {
					delete(differ, tempWord)
				}
			}
			// 如果map中一个元素都没有（因为频次差为 0 时，会执行 map delete操作），
			// 证明此时窗口中包含的正好是符合条件的一种组合
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}
