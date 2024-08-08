package jianzhi

/**
 * preorderTraversal
 *  @Description: 使用Morris方法实现二叉树前序遍历
 *  @param root
 *  @return []int
 */
func preorderTraversal(root *TreeNode) (result []int) {
	curNode := root
	for curNode != nil {
		leftNode := curNode.Left
		if leftNode == nil {
			result = append(result, curNode.Val)
			curNode = curNode.Right
			continue
		}
		for leftNode.Right != nil && leftNode.Right != curNode {
			leftNode = leftNode.Right
		}
		if leftNode.Right == nil {
			result = append(result, curNode.Val)
			leftNode.Right = curNode
			curNode = curNode.Left
		} else {
			leftNode.Right = nil
			curNode = curNode.Right
		}
	}
	return
}

/**
 * inorderTraversal
 *  @Description: 使用Morris方法实现二叉树中序遍历
 *  @param root
 *  @return result
 */
func inorderTraversal(root *TreeNode) (result []int) {
	curNode := root
	for curNode != nil {
		leftNode := curNode.Left
		if leftNode == nil {
			result = append(result, curNode.Val)
			curNode = curNode.Right
			continue
		}
		for leftNode.Right != nil && leftNode.Right != curNode {
			leftNode = leftNode.Right
		}
		if leftNode.Right == nil {
			leftNode.Right = curNode
			curNode = curNode.Left
		} else {
			result = append(result, curNode.Val)
			leftNode.Right = nil
			curNode = curNode.Right
		}
	}
	return
}

/**
 * postorderTraversal
 *  @Description: 使用Morris方法实现二叉树后序遍历
 *  @param root
 *  @return result
 */
func postorderTraversal(root *TreeNode) (result []int) {
	addPath := func(node *TreeNode) {
		size := len(result)
		for node != nil {
			result = append(result, node.Val)
			node = node.Right
		}
		// 因为是后序遍历，所以需要将本次添加的右孩子集列表反转一下，得到“左 右 根”的result
		reverse(result[size:])
	}
	curNode := root
	for curNode != nil {
		leftNode := curNode.Left
		if leftNode == nil {
			curNode = curNode.Right
			continue
		}
		for leftNode.Right != nil && leftNode.Right != curNode {
			leftNode = leftNode.Right
		}
		if leftNode.Right == nil {
			leftNode.Right = curNode
			curNode = curNode.Left
		} else {
			leftNode.Right = nil
			addPath(curNode.Left)
			curNode = curNode.Right
		}
	}
	addPath(root)
	return
}

func reverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
