package jianzhi

/**
 * rectCover
 *  @Description: 用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形
 *  @param num
 *  @return int
 */
func rectCover(num int) int {
	var third int
	if num < 4 {
		return num
	}

	first, second := 2, 3
	for i := 4; i <= num; i++ {
		// 取模
		third = second + first
		first = second
		second = third
	}
	return third
}
