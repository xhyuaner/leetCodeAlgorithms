package jianzhi

func spiralOrder(matrix [][]int) (result []int) {
	top, left := 0, 0
	bottom, right := len(matrix)-1, len(matrix[0])-1
	for left <= right && top <= bottom {
		direction := left
		for ; direction <= right; direction++ {
			result = append(result, matrix[top][direction])
		}
		direction = top + 1
		for ; direction <= bottom; direction++ {
			result = append(result, matrix[direction][right])
		}
		if left < right && top < bottom {
			direction = right - 1
			for ; direction >= left; direction-- {
				result = append(result, matrix[bottom][direction])
			}
			direction = bottom - 1
			for ; direction > top; direction-- {
				result = append(result, matrix[direction][left])
			}
		}
		top, left = top+1, left+1
		bottom, right = bottom-1, right-1
	}
	return result
}
