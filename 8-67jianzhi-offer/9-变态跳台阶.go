package main

import "math"

/**
 * btTiao
 *  @Description: 有 num 个台阶，每次可以选择跳 一个 或者 两个台阶 或者 n个台阶。求有多少种不同的跳跃方式。结果需要取模
 *  @param num
 *  @return int
 */
func btTiao(num int) int {
	/**
	第一步有n种跳法：跳1级、跳2级、到跳n级
		如果第一步跳1级，剩下n-1级，则剩下跳法是f(n-1)
		如果第一步跳2级，剩下n-2级，则剩下跳法是f(n-2)
	    。。。
	所以f(n)=f(n-1)+f(n-2)+...+f(1) 因为f(n-1)=f(n-2)+f(n-3)+...+f(1)
	所以f(n)=2*f(n-1)
	*/
	if num == 1 {
		return 1
	}
	return int(math.Pow(2, float64(num-1)))
}
