package jianzhi

var index int

func largestIsland(grid [][]int) int {
	maxArea := 0
	var areaArray []int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				index = len(areaArray) + 2
				area := dfsLand(grid, r, c)
				// 将单个最大岛屿的面积记录下来，避免无法填海或者填海后面积没有单个面积大时返回了0值
				maxArea = max(maxArea, area)
				// 获取每一个岛屿的面积，并添加到数组中
				areaArray = append(areaArray, area)
			}
		}
	}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				sumArea := 0
				for k := range landsIndex(grid, r, c) {
					// 从areaArray中获取当前岛屿的面积,-2是因为前面+2了
					sumArea += areaArray[k-2]
				}
				sumArea++ // 加上自身增加的那块面积
				maxArea = max(maxArea, sumArea)
			}
		}
	}
	return maxArea
}

func landsIndex(grid [][]int, r, c int) map[int]bool {
	indexs := make(map[int]bool)
	zuobiao := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, v := range zuobiao {
		row, col := r, c
		row += v[0]
		col += v[1]
		if inArea(grid, row, col) && grid[row][col] != 0 && indexs[grid[row][col]] == false {
			indexs[grid[row][col]] = true
		}
	}
	return indexs
}

func dfsLand(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) || grid[r][c] != 1 {
		return 0
	}
	// 将岛屿中每一块的数值标为大于1的索引，表示已经遍历过，并可通过这个索引（要-2）从areaArray中获取当前岛屿的面积
	grid[r][c] = index
	return 1 + dfsLand(grid, r-1, c) + dfsLand(grid, r+1, c) + dfsLand(grid, r, c-1) + dfsLand(grid, r, c+1)
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
