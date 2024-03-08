package jianzhi

func numIslands(grid [][]byte) (ans int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				ans++
				dfsLand(grid, r, c)
			}
		}
	}
	return ans
}

func dfsLand(grid [][]byte, r, c int) {
	if !inArea(grid, r, c) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2' // 2 表示已经遍历过的陆地
	dfsLand(grid, r-1, c)
	dfsLand(grid, r+1, c)
	dfsLand(grid, r, c-1)
	dfsLand(grid, r, c+1)
}

// 判断是否超出边界
func inArea(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
