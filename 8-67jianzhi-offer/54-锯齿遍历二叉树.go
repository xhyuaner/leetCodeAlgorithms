package jianzhi

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if len(ans)%2 > 0 {
				vals[n-1-i] = node.Val
			} else {
				vals[i] = node.Val
			}
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
