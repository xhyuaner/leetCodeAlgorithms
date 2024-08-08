package jianzhi

/*
*
  - leastInterval
  - @Description: 给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表，用字母 A 到 Z 表示，
    以及一个冷却时间 n。每个周期或时间间隔允许完成一项任务。任务可以按任何顺序完成，
    但有一个限制：两个 相同种类 的任务之间必须有长度为 n 的冷却时间。
    返回完成所有任务所需要的 最短时间间隔 。
  - @param tasks
  - @param n
  - @return int

输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B

	在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
*/
func leastInterval(tasks []byte, n int) int {
	var maxCount, maxCountKinds int
	taskArray := [26]int{}
	for _, v := range tasks {
		taskArray[v-'A']++
		maxCount = max(maxCount, taskArray[v-'A'])
	}
	for _, v := range taskArray {
		if v == maxCount {
			maxCountKinds++
		}
	}
	result := (maxCount-1)*(n+1) + maxCountKinds
	// 取最大是为了避免n=0的情况
	return max(result, len(tasks))
}
