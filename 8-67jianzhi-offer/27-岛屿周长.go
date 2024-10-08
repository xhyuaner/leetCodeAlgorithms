package jianzhi

/**
 * islandPerimeter
 *  @Description: 给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。
				网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，
				但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
				岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，
				且宽度和高度均不超过 100 。计算这个岛屿的周长。
 *  @param grid
 *  @return int
*/
func islandPerimeter(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				return dfsLand2(grid, r, c)
			}
		}
	}
	return 0
}

func dfsLand2(grid [][]int, r, c int) int {
	if !inArea2(grid, r, c) {
		return 1
	}
	if grid[r][c] == 0 {
		return 1
	}
	if grid[r][c] == 2 {
		return 0
	}
	grid[r][c] = 2
	return dfsLand2(grid, r-1, c) + dfsLand2(grid, r+1, c) + dfsLand2(grid, r, c-1) + dfsLand2(grid, r, c+1)
}

func inArea2(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
