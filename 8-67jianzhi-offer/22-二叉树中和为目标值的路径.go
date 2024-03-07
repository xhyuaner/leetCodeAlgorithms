package jianzhi

func pathTarget(root *TreeNode, target int) [][]int {
	var ans [][]int
	var dfs func(root *TreeNode, cur int, list []int)
	dfs = func(root *TreeNode, cur int, list []int) {
		if root == nil {
			return
		}
		cur += root.Val
		list = append(list, root.Val)
		if cur == target && root.Left == nil && root.Right == nil {
			// 如果这里这样写，由于go是值传递（会将变量拷贝一份再作为参数传下去），这里的list其实是个指向底层数组的指针
			// 这个指针会被拷贝继续往下传，但是经过append后指针指向的底层数组里的元素会改变
			// 所以就可能导致最终ans发生改变，因为ans中存的是切片，切片不变，但是切片里指向的底层数组发生了改变
			// ans = append(ans, list)
			// 为了解决上面的问题，我们将list切片里面的元素拷贝到一个新切片中，放入ans，这个新切片不会再被改变
			ans = append(ans, append([]int(nil), list...))
			return
		}
		dfs(root.Left, cur, list)
		dfs(root.Right, cur, list)
		// 可不写，go的参数传递是值传递(值拷贝传递), 而引用类型的例如slice会随着递归的回溯过程返回上一层保存的副本。
		// 回溯过程中的"破坏现场"“恢复现场”的操作甚至不需要显示的去做。就可以保证回溯的效果。
		// list = list[:len(list)-1]
	}
	dfs(root, 0, []int{})
	return ans
}
