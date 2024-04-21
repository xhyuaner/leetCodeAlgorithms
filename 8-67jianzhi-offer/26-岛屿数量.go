package jianzhi

/**
 * numIslands
 *  @Description: 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
				岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
				此外，你可以假设该网格的四条边均被水包围。
 *  @param grid
 *  @return ans
*/
func numIslands(grid [][]byte) (ans int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				ans++
				dfsLand1(grid, r, c)
			}
		}
	}
	return ans
}

func dfsLand1(grid [][]byte, r, c int) {
	if !inArea1(grid, r, c) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2' // 2 表示已经遍历过的陆地
	dfsLand1(grid, r-1, c)
	dfsLand1(grid, r+1, c)
	dfsLand1(grid, r, c-1)
	dfsLand1(grid, r, c+1)
}

// 判断是否超出边界
func inArea1(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
