package jianzhi

/**
 * numberOf1
 *  @Description: 输入一个整数，输出该数32位二进制表示中1的个数。其中负数用补码表示
 *  @param num
 *  @return int
 */
func numberOf1(num int) int {
	count := 0
	for num != 0 {
		count++
		// 通过 与 运算消掉二进制最右边那个1, 看总共可以消多少次，就有多少个1
		// 例如：1100 & 1011 = 1000
		num &= num - 1
	}
	return count
}
