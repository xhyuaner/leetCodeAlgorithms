package jianzhi

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
*/

func coinChange(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, amount+1)
		for j := range f[i] {
			f[i][j] = 99999
		}
	}
	f[0][0] = 0
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if x > c {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-x]+1) // 因为可以重复选，所以还是“f[(i+1)%2][c-x]+1”
			}
		}
	}
	ans := f[n%2][amount]
	if ans < 9999 {
		return ans
	} else {
		return -1
	}
}

func coinChange2(coins []int, amount int) int {
	//n := len(coins)
	f := make([]int, amount+1)
	for i := range f {
		f[i] = 999999
	}
	f[0] = 0
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			f[c] = min(f[c], f[c-x]+1) // 因为可以重复选，所以还是“f[(i+1)%2][c-x]+1”
		}
	}
	ans := f[amount]
	if ans < 999999 {
		return ans
	} else {
		return -1
	}
}
