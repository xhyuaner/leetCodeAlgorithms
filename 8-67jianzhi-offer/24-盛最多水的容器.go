package jianzhi

/**
 * maxArea
 *  @Description: 盛最多水的容器---接雨水基础版
 *  @param height
 *  @return int
 */
func maxArea(height []int) (maxArea int) {
	l, r := 0, len(height)-1
	for l < r {
		maxArea = max(maxArea, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return
}
