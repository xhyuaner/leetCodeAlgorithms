package jianzhi

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
