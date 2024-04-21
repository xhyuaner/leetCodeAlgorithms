package jianzhi

/**
给定一个正整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
*/

/**
 * findTargetSumWays
 *  @Description: 两个数组（滚动数组）
 *  @param nums
 *  @param target
 *  @return int
 */
func findTargetSumWays(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	n := len(nums)
	arr := [2][]int{make([]int, target+1), make([]int, target+1)}
	arr[0][0] = 1
	for i, x := range nums {
		for c := 0; c <= target; c++ {
			if x > c {
				arr[(i+1)%2][c] = arr[i%2][c]
			} else {
				arr[(i+1)%2][c] = arr[i%2][c] + arr[i%2][c-x]
			}
		}

	}
	return arr[n%2][target]
}

/**
 * findTargetSumWays2
 *  @Description: 一个数组解决
 *  @param nums
 *  @param target
 *  @return int
 */
func findTargetSumWays2(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	//n := len(nums)
	arr := make([]int, target+1)
	arr[0] = 1
	for _, x := range nums {
		for c := target; c >= x; c-- {
			arr[c] += arr[c-x]
		}

	}
	return arr[target]
}
