package main

import "fmt"

type UnionFindSet struct {
	parent []int
}

func NewUnionFindSet(size int) *UnionFindSet {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFindSet{parent}
}

func (uf *UnionFindSet) find(x int) int {
	for uf.parent[x] != x {
		x = uf.parent[x] // 此题不适合做路径压缩
		//uf.parent[x] = uf.find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFindSet) join(x, y int) {
	if uf.find(x) != uf.find(y) {
		uf.parent[uf.find(x)] = uf.find(y)
	}
}

func (uf *UnionFindSet) IsConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func main() {
	var n, m, p int
	fmt.Scanln(&n, &m, &p)
	union := NewUnionFindSet(n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		union.join(x-1, y-1)
	}
	for i := 0; i < p; i++ {
		var op, x, y int
		fmt.Scanln(&op, &x, &y)
		if op == 1 {
			if union.find(x-1) == union.find(y-1) {
				union.parent[x-1] = x - 1
			}
		}
		if op == 2 {
			if union.IsConnected(x-1, y-1) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
