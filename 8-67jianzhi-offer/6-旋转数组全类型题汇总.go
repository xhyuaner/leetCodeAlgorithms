package jianzhi

import "fmt"

func main() {
	nums := []int{21, 21, -21, -20, -17, -8, -6, -2, -2, -1, 0, 2, 3, 4, 4, 6, 11, 13, 14, 16, 17, 18, 20}
	fmt.Println(search3(nums, 4))
}

/**
 * search3
 *  @Description:从含有重复值的旋转数组中找到目标值，若有多个相同元素，返回索引值最小的一个。
 *  @param nums
 *  @param target
 *  @return int
 */
func search3(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[left] == target {
			return left
		}
		if nums[mid] == target {
			if mid-left-1 < 0 || nums[mid-1] != nums[mid] {
				return mid
			}
			right = mid - 1
			continue
		}
		// 左半部分有序
		if nums[mid] > nums[right] {
			// 如果target在左边有序区域
			if nums[left] < target && target < nums[mid] {
				right = mid - 1
			} else { // 否则
				left = mid + 1
			}
		} else if nums[mid] < nums[right] { // 右半部分有序
			// 如果target在右边有序区域
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else { // 否则
				right = mid - 1
			}
		} else {
			right--
		}
	}
	return -1
}

/**
 * findMin1
 *  @Description: 从不含有重复值的旋转数组中查找最小值
 *  @param nums
 *  @return int
 */
func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

/**
 * findMin2
 *  @Description: 从含有重复值的旋转数组中查找最小值
 *  @param nums
 */
func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

/**
 * search1
 *  @Description: 从不含有重复值的旋转数组中查找目标值
 *  @param nums
 */
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 左半部分有序
		if nums[mid] > nums[right] {
			// 如果target在左边有序区域
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else { // 否则
				left = mid + 1
			}
		} else { // 右半部分有序
			// 如果target在右边有序区域
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else { // 否则
				right = mid - 1
			}
		}
	}
	return -1
}

/**
 * search2
 *  @Description: 从含有重复值的旋转数组中查找是否含有目标值
 *  @param nums
 */
func search2(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		// 左半部分有序
		if nums[mid] > nums[right] {
			// 如果target在左边有序区域
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else { // 否则
				left = mid + 1
			}
		} else if nums[mid] < nums[right] { // 右半部分有序
			// 如果target在右边有序区域
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else { // 否则
				right = mid - 1
			}
		} else {
			right--
		}
	}
	return false
}
