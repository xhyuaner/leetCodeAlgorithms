package sort_doublePointer

import "sort"

/**
Tag：排序 + 双指针

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

时间复杂度：O(n^2)
空间复杂度：O(logN)
*/
func threeSum(nums []int) [][]int {
	var ans [][]int
	// 对数组进行排序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 如果第一个固定数大于0，则不存在和为0的情况，直接返回
		if nums[i] > 0 {
			return ans
		}
		// 如果当前固定数和前一个固定数相同，则整个数组已在之前被遍历过一次，直接continue
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 定义左指针和右指针
		left, right := i+1, len(nums)-1

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去除left重复值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去除right重复值
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
		}
	}
	return ans
}
