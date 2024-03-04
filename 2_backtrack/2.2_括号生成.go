package backtrack

/**
Tag: 回溯算法

Description:
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

Example:
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

Analysis:
时间复杂度：O(4^n / n^(1/2))
空间复杂度：O(n)
*/
func generateParenthesis(n int) (ans []string) {
	// 定义一个临时字符数组
	var cur []rune
	// 递归生成括号组合
	parenthesesBackTrack(&ans, &cur, 0, 0, n)
	return ans
}

func parenthesesBackTrack(ans *[]string, cur *[]rune, open, close, max int) {
	// 递归出口
	if len(*cur) == 2*max { // 如果当前字符数组的长度达到 2n，就添加到结果集
		*ans = append(*ans, string(*cur))
	}
	// 在添加括号之前进行合法性判断
	// "("数量必须小于 n
	if open < max {
		*cur = append(*cur, '(')
		// 递归调用
		parenthesesBackTrack(ans, cur, open+1, close, max)
		*cur = (*cur)[:len(*cur)-1]
	}
	// 同时 ")" 数量必须小于当前字符数组中已有的 ”(“ 数量
	if close < open {
		*cur = append(*cur, ')')
		// 递归调用
		parenthesesBackTrack(ans, cur, open, close+1, max)
		// 当有一种符合要求的排列组合被添加到ans并返回，就删除当前字符数组的最后一位，递归出栈，探索下一种组合
		*cur = (*cur)[:len(*cur)-1]
	}
}
