package binary_search

/**
Tag: 二分查找

Description:
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

Example:
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：
输入：nums = [1], target = 0
输出：-1

Analysis:
时间复杂度：O(log n)，其中 n 为 nums 数组的大小。整个算法时间复杂度即为二分查找的时间复杂度 O(log n)。
空间复杂度：O(1)
*/

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	var mid int
	left, right := 0, n-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 中间值左边属于有序部分
			if nums[0] <= target && target < nums[mid] { // target在mid左边，正常二分查找
				right = mid - 1
			} else {
				left = mid + 1 // target在mid右边
			}
		} else { // 中间值右边属于有序部分
			if nums[mid] < target && target <= nums[n-1] { // target在mid右边，正常二分查找
				left = mid + 1
			} else {
				right = mid - 1 // target在mid左边
			}
		}
	}
	return -1
}
