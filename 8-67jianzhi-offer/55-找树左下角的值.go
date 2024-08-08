package jianzhi

func findBottomLeftValue(root *TreeNode) int {
	node := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}
