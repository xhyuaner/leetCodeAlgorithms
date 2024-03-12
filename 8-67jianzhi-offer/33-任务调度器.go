package jianzhi

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
	return max(result, len(tasks))
}
