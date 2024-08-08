package jianzhi

/*
*
  - maxSubArray
  - @Description:
    输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
    -2,1,-2,4,3,5,6,1,5
    输出：6
    解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
  - @param nums
  - @return int
*/
func MaxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
