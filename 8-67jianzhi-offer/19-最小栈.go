package jianzhi

// 请你设计一个 最小栈 。它提供 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	a []int // 数据栈，存储所有元素
	b []int // 辅助栈，存储a中非严格降序元素

}

func Constructor2() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.a = append(this.a, x)
	if len(this.b) == 0 || x <= this.b[len(this.b)-1] {
		this.b = append(this.b, x)
	}
}

func (this *MinStack) Pop() {
	if this.a[len(this.a)-1] == this.b[len(this.b)-1] {
		this.b = this.b[:len(this.b)-1]
	}
	this.a = this.a[:len(this.a)-1]
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MinStack) GetMin() int {
	return this.b[len(this.b)-1]
}
