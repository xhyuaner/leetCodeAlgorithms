package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用map存储中序遍历列表，加速树节点Val的查找时间
var inorderMap map[int]int
var preord []int

/**
 * deduceTree
 *  @Description:
 *  @param preorder
 *  @param inorder
 *  @return *TreeNode
 */
func deduceTree(preorder []int, inorder []int) *TreeNode {
	preord = preorder
	inorderMap = make(map[int]int, len(inorder))
	// 将中序遍历列表存入map中
	for i, v := range inorder {
		inorderMap[v] = i
	}
	// 返回递归子树
	return recur(0, 0, len(inorder)-1)
}

/**
 * recur
 *  @Description:递归函数
 *  @param rootIndex
 *  @param left
 *  @param right
 *  @return *TreeNode
 */
func recur(rootIndex, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	root := preord[rootIndex]
	newTree := &TreeNode{Val: root}
	i := inorderMap[root]
	newTree.Left = recur(rootIndex+1, left, i-1)
	newTree.Right = recur(rootIndex+i-left+1, i+1, right)
	return newTree
}
