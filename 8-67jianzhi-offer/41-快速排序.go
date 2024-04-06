package jianzhi

func sortArray(nums []int) []int {
	qSort(nums, 0, len(nums)-1)
	return nums
}
func qSort(nums []int, l, r int) {
	if l < r {
		p := partition1(nums, l, r)
		qSort(nums, l, p-1)
		qSort(nums, p+1, r)
	}
}
func partition1(nums []int, l, r int) int {
	mid := l + (r-l)/2
	p := nums[mid]
	nums[mid], nums[r] = nums[r], nums[mid]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
