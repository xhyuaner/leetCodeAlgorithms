package jianzhi

func islandPerimeter(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				return dfsLand(grid, r, c)
			}
		}
	}
	return 0
}

func dfsLand(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 1
	}
	if grid[r][c] == 0 {
		return 1
	}
	if grid[r][c] == 2 {
		return 0
	}
	grid[r][c] = 2
	return dfsLand(grid, r-1, c) + dfsLand(grid, r+1, c) + dfsLand(grid, r, c-1) + dfsLand(grid, r, c+1)
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
