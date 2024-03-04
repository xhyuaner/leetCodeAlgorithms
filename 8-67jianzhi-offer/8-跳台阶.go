package main

/**
 * trainWays
 *  @Description: 有 num 个台阶，每次可以选择跳 一个 或者 两个台阶。求有多少种不同的跳跃方式。结果需要取模
 *  @param num
 *  @return int
 */
func trainWays(num int) int {
	const mod int = 1e9 + 7
	var third int
	// 如果没有台阶，就不跳，假如不跳也算一种方式
	if num == 0 {
		return 1
	}

	if num < 3 {
		return num
	}
	first, second := 1, 2
	for i := 3; i <= num; i++ {
		// 取模
		third = (second + first) % mod
		first = second
		second = third
	}
	return third
}
