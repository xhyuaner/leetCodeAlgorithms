package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var result []int
	top, bottom := 0, len(matrix)-1
	left := 0
	dir := 0
	// 针对二维数组中各一维数组长度不一样的情况，设置一个右边界索引map，存储每行的右边界
	rightMap := make(map[int]int, len(matrix))
	for i, arr := range matrix {
		rightMap[i] = len(arr) - 1
	}

	for top <= bottom {
		if dir == 0 { // 向右移动
			for i := left; i <= rightMap[top] && top <= bottom; i++ {
				result = append(result, matrix[top][i])
			}
			top++
		} else if dir == 1 { // 向下移动
			for i := top; i <= bottom && left <= rightMap[i]; i++ {
				result = append(result, matrix[i][rightMap[i]])
				rightMap[i]-- // 向下遍历的过程中，同时将右边界减一
			}
		} else if dir == 2 { // 向左移动
			for i := rightMap[bottom]; i >= left && top <= bottom; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		} else if dir == 3 { // 向上移动
			for i := bottom; i >= top && left <= rightMap[i]; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
		dir = (dir + 1) % 4
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 4},
		{5, 6, 7, 99},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	fmt.Println(spiralOrder(matrix))
}
