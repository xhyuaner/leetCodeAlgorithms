package jianzhi

/**
 * findTargetIn2DPlants
 *  @Description: 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
					每一列都按照从上到下递增的顺序排序。
					请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 *  @param plants
 *  @param target
 *  @return bool
*/
func findTargetIn2DPlants(plants [][]int, target int) bool {
	// 将二维数组旋转45°，看成二叉搜索树，从左下角开始遍历
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
