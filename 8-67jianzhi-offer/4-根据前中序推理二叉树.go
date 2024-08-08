package jianzhi

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * deduceTree
 *  @Description:
 *  @param preorder
 *  @param inorder
 *  @return *TreeNode
 */
func deduceTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	// 将中序遍历列表存入map中
	// 使用map存储中序遍历列表，加速树节点Val的查找时间
	for i, v := range inorder {
		inorderMap[v] = i
	}
	var recur func(rootIndex, left, right int) *TreeNode
	recur = func(rootIndex, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		root := preorder[rootIndex]
		newTree := &TreeNode{Val: root}
		i := inorderMap[root]
		newTree.Left = recur(rootIndex+1, left, i-1)
		// rootIndex+(i-left)+1：根节点索引 + 左子树长度 + 1 = 右子树的根结点索引
		newTree.Right = recur(rootIndex+(i-left)+1, i+1, right)
		return newTree
	}
	// 返回递归子树
	return recur(0, 0, len(inorder)-1)
}
