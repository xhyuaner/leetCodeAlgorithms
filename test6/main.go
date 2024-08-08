package main

import "math"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func isWQtree(t *tree) bool {
	var queue []*tree
	flag := 0
	queue = append(queue, t)
	for len(queue) != 0 {
		t0 := queue[0]
		if t0.left != nil {
			if flag == 1 {
				return false
			}
			queue = append(queue, t0.left)
		} else {
			flag = 1
		}
		if t0.right != nil {
			if flag == 1 {
				return false
			}
			queue = append(queue, t0.right)
		} else {
			flag = 1
		}
		queue = queue[1:]
	}

	return true
	//
	//if t!=nil{
	//	return isWQtree(t.left)
	//}

}

func isAVGtree(t *tree) (bool, int) {
	//if t==nil {
	//	return false
	//}
	//var dfs func(*tree) int
	if t == nil {
		return true, 0
	}
	isL, levelL := isAVGtree(t.left)
	if !isL {
		return false, 0
	}
	isR, levelR := isAVGtree(t.right)
	if !isR {
		return false, 0
	}
	if math.Abs(float64(levelL-levelR)) <= 1 {
		return true, max(levelL, levelR) + 1
	}
	return false, 0

	//return max(dfs(t.left)+1,dfs(t.right)+1)
	//leftLevel,rightLevel := dfs(t.left),dfs(t.right)
	//if isAVGtree(t.left)&&isAVGtree(t.right){
	//	if leftLevel-rightLevel<=1&&rightLevel-leftLevel<=1 {
	//		return true
	//	}
	//}
	//return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
