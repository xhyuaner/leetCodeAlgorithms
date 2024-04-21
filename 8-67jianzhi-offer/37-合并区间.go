package jianzhi

import "sort"

/**
 * merge
 *  @Description: 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
				输出：[[1,6],[8,10],[15,18]]
				解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *  @param intervals
 *  @return [][]int
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}
