package jianzhi

/**
 * trap
 *  @Description: 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *  @param height
 *  @return int
 */
func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	preMax, sufMax := 0, 0
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}
