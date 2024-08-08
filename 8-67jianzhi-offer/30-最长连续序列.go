package jianzhi

/*
*
  - longestConsecutive
  - @Description: 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
  - @param nums
  - @return int

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
时间复杂度：O(n)
空间复杂度：O(n)
*/
func longestConsecutive(nums []int) (ans int) {
	numsMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numsMap[num] = true
	}
	for key := range numsMap {
		if !numsMap[key-1] {
			curNum := key
			maxLen := 1
			for numsMap[curNum+1] {
				maxLen++
				curNum++
			}
			ans = max(ans, maxLen)
		}
	}
	return
}
