package jianzhi

/**
 * isSubStructure
 *  @Description: 给定两棵二叉树 tree1 和 tree2，判断 tree2 是否以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
					注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值
 *  @param A
 *  @param B
 *  @return bool
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if isSame(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}
