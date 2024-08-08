package jianzhi

/**
 * trap
 *  @Description: 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *  @param height
 *  @return int
 */
func trap(height []int) (ans int) {
	l, r := 0, len(height)-1
	var preMax, sufMax int
	for l <= r {
		preMax = max(preMax, height[l])
		sufMax = max(sufMax, height[r])
		if preMax < sufMax {
			ans += preMax - height[l]
			l++
		} else {
			ans += sufMax - height[r]
			r--
		}
	}
	return
}
