package sort_doublePointer

/**
Tag: 倒排

Description:
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，
那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。

Example:
输入：nums = [1,2,3]
输出：[1,3,2]
输入：nums = [3,2,1]
输出：[1,2,3]
输入：nums = [1,1,5]
输出：[1,5,1]

Analysis:
时间复杂度：O(n)
空间复杂度：O(1)
*/
func nextPermutation(nums []int) []int {
	n := len(nums)
	var i, left int
	small := -1
	for i = n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// 确定较小值索引
			small = i - 1
			break
		}
	}
	if small != -1 { // 找到较小值
		// 寻找右边的较大值的下一个索引
		for i < len(nums) && nums[i] > nums[small] {
			i++
		}
		// 交换较小值和右边较大值
		nums[small], nums[i-1] = nums[i-1], nums[small]
		// 利用双指针倒序较小值右边部分的数组
		left = small + 1
	} else { // 没找到较小值（数组整体呈现降序排列），就让倒排左指针指向第一个元素，数组整体倒排
		left = 0
	}
	// 倒排数组
	for right := n - 1; right > left; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
