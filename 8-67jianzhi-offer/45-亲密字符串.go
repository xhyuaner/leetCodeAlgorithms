package jianzhi

/**
 * buddyStrings
 *  @Description:给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，
					就返回 true ；否则返回 false 。
 *  @param s
 *  @param goal
 *  @return bool
*/
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		se := [26]bool{}
		for _, ch := range s {
			if se[ch-'a'] {
				return true
			}
			se[ch-'a'] = true
		}
		return false
	}
	first, second := -1, -1
	for i, _ := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
