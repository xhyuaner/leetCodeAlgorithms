package binary_search

/**
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
	leftIdx := binarySearch(nums, target, true)
	rightIdx := binarySearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		return []int{leftIdx, rightIdx}
	}
	return []int{-1, -1}

}
func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}
