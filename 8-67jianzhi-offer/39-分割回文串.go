package jianzhi

/**
 * partition
 *  @Description: 给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。
 *  @param s
 *  @return ans
*/
func partition(s string) (ans [][]string) {
	var path []string
	n := len(s)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		for j := i; j < n; j++ {
			if isHWC(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func isHWC(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
