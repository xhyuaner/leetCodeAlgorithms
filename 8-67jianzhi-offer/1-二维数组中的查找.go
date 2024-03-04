package main

import "fmt"

func main() {
	list := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}}
	fmt.Println(findTargetIn2DPlants(list, 11))
}

// 将二维数组旋转45°，看成二叉搜索树，从左下角开始遍历
func findTargetIn2DPlants(plants [][]int, target int) bool {
	i, j := len(plants)-1, 0
	for i >= 0 && j < len(plants[0]) {
		if plants[i][j] == target {
			return true
		} else if plants[i][j] < target {
			j++
		} else if plants[i][j] > target {
			i--
		}
	}
	return false
}
