package jianzhi

/*
*
  - validateStackSequences
  - @Description:	给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，
    只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。
  - @param pushed
  - @param popped
  - @return bool
    输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
    输出：true
    解释：我们可以按以下顺序执行：
    push(1), push(2), push(3), push(4), pop() -> 4,
    push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。
*/
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

func validateStackSequences2(pushed []int, popped []int) bool {
	var stk []int
	idx := 0
	for _, v := range pushed {
		stk = append(stk, v)
		for len(stk) != 0 && stk[len(stk)-1] == popped[idx] {
			idx++
			stk = stk[0 : len(stk)-1]
		}
	}
	return len(stk) == 0
}
