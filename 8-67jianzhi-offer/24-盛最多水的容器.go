package jianzhi

/**
 * maxArea
 *  @Description: 盛最多水的容器---接雨水基础版
 *  @param height
 *  @return int
 */
func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
