package jianzhi

/**
 * leastInterval
 *  @Description: 给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表，用字母 A 到 Z 表示，
					以及一个冷却时间 n。每个周期或时间间隔允许完成一项任务。任务可以按任何顺序完成，
					但有一个限制：两个 相同种类 的任务之间必须有长度为 n 的冷却时间。
					返回完成所有任务所需要的 最短时间间隔 。
 *  @param tasks
 *  @param n
 *  @return int
*/
func leastInterval(tasks []byte, n int) int {
	var maxCount, maxKinds int
	taskArray := [26]int{}
	for _, v := range tasks {
		taskArray[v-'A']++
		maxCount = max(maxCount, taskArray[v-'A'])
	}
	for _, v := range taskArray {
		if v == maxCount {
			maxKinds++
		}
	}
	result := (maxCount-1)*(n+1) + maxKinds
	// 取最大是为了避免n=0的情况
	return max(result, len(tasks))
}
