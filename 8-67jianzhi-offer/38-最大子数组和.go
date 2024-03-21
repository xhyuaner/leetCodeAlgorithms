package jianzhi

import "math"

func maxSubArray(nums []int) int {
	ans := math.MinInt
	// [i,j]子数组和 = [0,j]前缀和 - [0,i]前缀和
	// 要让子数组和最大，就要让 [0,i]前缀和 最小
	// 所以维护了 当前前缀和preSum 和 最小前缀和minPreSum
	minPreSum, preSum := 0, 0
	for _, num := range nums {
		preSum += num
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
