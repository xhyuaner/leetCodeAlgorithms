package main

/**
给定一个长度为 n 的整数数组 height 。
有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxArea(height []int) int {
	MaxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		MaxArea = max(MaxArea, area)
		// 能组成更大面积的两条边必定 不包含 此刻这条较短的边，所以哪条短就淘汰哪条继续找能组成更大面积的边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return MaxArea
}
