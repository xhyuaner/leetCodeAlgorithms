package jianzhi

/**
 * preorderTraversal
 *  @Description: 使用Morris方法实现二叉树前序遍历
 *  @param root
 *  @return []int
 */
func preorderTraversal(root *TreeNode) (result []int) {
	p1 := root
	for p1 != nil {
		p2 := p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				result = append(result, p1.Val)
				p2.Right = p1
				p1 = p1.Left
			} else {
				p2.Right = nil
				p1 = p1.Right
			}
		} else {
			result = append(result, p1.Val)
			p1 = p1.Right
		}
	}
	return result
}

/**
 * inorderTraversal
 *  @Description: 使用Morris方法实现二叉树中序遍历
 *  @param root
 *  @return result
 */
func inorderTraversal(root *TreeNode) (result []int) {
	p1 := root
	for p1 != nil {
		p2 := p1.Left
		if p2 != nil {
			// 循环查找p1左子树的最右边结点
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil { // 将左子树最右边结点的Right指针指向p1
				p2.Right = p1
				p1 = p1.Left
			} else {
				result = append(result, p1.Val)
				// p1回到上次所在位置后，再次寻找左子树最右边结点，发现已经有了连线，此时切断连线，并将p1.Left添加到返回列表
				p2.Right = nil
				p1 = p1.Right
			}
		} else {
			result = append(result, p1.Val)
			// 如果p1的左子树为空，p1直接去指向p1的Right，或者是通过我们添加的最右结点连线回到p1上一次所在的位置
			p1 = p1.Right
		}
	}
	return
}

func reverse(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
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

	p1 := root
	for p1 != nil {
		p2 := p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
			} else {
				p2.Right = nil
				addPath(p1.Left)
				p1 = p1.Right
			}
		} else {
			p1 = p1.Right
		}
	}
	addPath(root)
	return
}
