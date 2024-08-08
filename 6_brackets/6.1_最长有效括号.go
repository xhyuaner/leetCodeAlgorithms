package brackets

/*
*
Tag: 倒排

Description:
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

Example:
输入：s = "(()"
输出：2
输入：s = ")()())"
输出：4
输入：s = ""
输出：0

Analysis:
时间复杂度：O(n)，其中 n 为字符串长度。我们只要正反遍历两边字符串即可。
空间复杂度：O(1)
*/
func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	// 遍历第一次，但无法统计“(()”这类情况（因为此时有效括号长度为 2 ,但是没有计算到）
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		// 如果左括号和有括号数量相等，就计算此时的有效括号数量，并更新最长有效括号数量
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left { // 从左往右遍历，右括号数大于左括号数（已无效），就作废（left和right值归0）
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	// 遍历第二次，但无法统计“())”这类情况（因为此时有效括号长度为 2 ,但是没有计算到）
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	// 两次遍历完，所有情况都考虑到，返回最长有效括号数量即可
	return maxLength
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 用栈：时间和空间复杂度：O(n)
func longestValidParentheses2(s string) int {
	maxAns := 0
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

// 动态规划:	...()	...))
func longestValidParentheses3(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}
