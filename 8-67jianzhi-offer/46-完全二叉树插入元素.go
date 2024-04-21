package jianzhi

//type TreeNode struct {
//	Val         int
//	Left, Right *TreeNode
//}
/**
完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。
设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。
使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
CBTInserter.get_root() 将返回树的根节点。

输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
输出：[null,3,4,[1,2,3,4,5,6,7,8]]
*/

type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor3(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	var candidate []*TreeNode
	var node *TreeNode
	for len(q) != 0 {
		node = q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root: root, candidate: candidate}
}

func (this *CBTInserter) Insert(v int) int {
	child := &TreeNode{Val: v}
	can := this.candidate[0]
	if can.Left == nil {
		can.Left = child
	} else {
		can.Right = child
		this.candidate = this.candidate[1:]
	}
	this.candidate = append(this.candidate, child)
	return can.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
