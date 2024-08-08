package jianzhi

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n) // 预分配空间
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return
}
