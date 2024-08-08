package binary_search

/*
*
Tag: 二分查找

Description:
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

Example:
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

Analysis:
时间复杂度：O(log n)，其中 n 为 nums 数组的大小。二分查找的时间复杂度为 O(log n)，一共会执行两次，因此总时间复杂度为 O(log n)。
空间复杂度：O(1)
*/
func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target) // 使用其中一种写法即可
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1} // nums 中没有 target
	}
	// 如果 start 存在，那么 end 必定存在
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

// lowerBound 返回最小的满足 nums[i] >= target 的 i
// 如果数组为空，或者所有数都 < target，则返回 len(nums)
// 要求 nums 是非递减的，即 nums[i] <= nums[i + 1]

// 闭区间写法
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 闭区间 [left, right]
	for left <= right {           // 区间不为空
		// 循环不变量：
		// nums[left-1] < target
		// nums[right+1] >= target
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // 范围缩小到 [mid+1, right]
		} else {
			right = mid - 1 // 范围缩小到 [left, mid-1]
		}
	}
	return left
}
