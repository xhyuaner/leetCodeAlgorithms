package jianzhi

/**
 * minSubArrayLen（滑动窗口）
 *  @Description: 给定一个含有 n 个正整数的数组和一个正整数 target 。
				找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组
 				[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *  @param target
 *  @param nums
 *  @return int
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sum, left := 0, 0
	ans := n + 1
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum-nums[left] >= target {
			sum -= nums[left]
			left++
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}
