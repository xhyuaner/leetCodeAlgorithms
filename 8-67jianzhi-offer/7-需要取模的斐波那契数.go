package jianzhi

/**
 * fib
 *  @Description: 动态规划，使用三个变量循环交换计算结果
 *  @param n
 *  @return int
 */
func fib(n int) int {
	const mod int = 1e9 + 7
	var third int
	if n < 2 {
		return n
	}
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		// 题目要求对结果取模，避免结果过大
		third = (second + first) % mod
		first = second
		second = third
	}
	return third
}
