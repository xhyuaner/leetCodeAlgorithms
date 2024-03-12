package jianzhi

import "math"

func ternaryToDecimal(ternary string) int {
	decimal := 0
	for i := 0; i < len(ternary); i++ {
		// 将每一位三进制数转为十进制数并加到结果上
		digit := int(ternary[i] - '0')
		decimal += digit * int(math.Pow(3, float64(len(ternary)-i-1)))
	}
	return decimal
}
