package sort_doublePointer

import (
	"sort"
)

/**
Tag：排序 + 双指针

给你一个长度为 n 的整数数组 nums 和 一个目标值 target。
请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。

时间复杂度：O(n^2)
空间复杂度：O(logN)
*/
//func main() {
//	nums := []int{-1000, -5, -5, -5, -5, -5, -5, -1, -1, -1}
//	threeSumClosest(nums, -14)
//}
func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	// 初始化结果值
	ans := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		// 如果当前固定数和前一个固定数相同，则整个数组已在之前被遍历过一次，直接continue
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < abs(target-ans) {
				ans = sum
			}
			if sum > target {
				// 去除right重复值
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if sum < target {
				// 去除left重复值
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else { // 如果三个数和与目标值相等，那已经是最接近的了，直接返回
				return ans
			}
		}
	}
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
