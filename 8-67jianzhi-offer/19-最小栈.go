package jianzhi

// MinStack 请你设计一个 最小栈 。它提供 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	dataStack      []int // 数据栈，存储所有元素
	auxiliaryStack []int // 辅助栈，存储a中非严格降序元素
}

func Constructor2() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.dataStack = append(s.dataStack, x)
	if len(s.auxiliaryStack) == 0 || x <= s.auxiliaryStack[len(s.auxiliaryStack)-1] {
		s.auxiliaryStack = append(s.auxiliaryStack, x)
	}
}

func (s *MinStack) Pop() {
	if s.dataStack[len(s.dataStack)-1] == s.auxiliaryStack[len(s.auxiliaryStack)-1] {
		s.auxiliaryStack = s.auxiliaryStack[:len(s.auxiliaryStack)-1]
	}
	s.dataStack = s.dataStack[:len(s.dataStack)-1]
}

func (s *MinStack) Top() int {
	return s.dataStack[len(s.dataStack)-1]
}

func (s *MinStack) GetMin() int {
	return s.auxiliaryStack[len(s.auxiliaryStack)-1]
}
