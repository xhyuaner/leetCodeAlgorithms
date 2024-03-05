package jianzhi

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for push, pop := 0, 0; push < len(pushed) && pop < len(popped); {
		stack = append(stack, pushed[push])
		push++
		for len(stack) > 0 && stack[len(stack)-1] == popped[pop] {
			stack = stack[:len(stack)-1]
			pop++
		}
	}
	return len(stack) == 0
}
