package main

import "fmt"

/**
 * main
 *  @Description:
一天，牛牛对一份字母文件执行了XOR运算!
文件中包含字符串S。要对字母进行XOR操作时，首先需要一个密钥。如果S= "ABCD”且key = 10，则执行XOR操作,字符串变为 “KHIN".
A的ASCII码为65,如果与10进行XOR运算，结果为75.由于75是ASCII码中的K，A被加密为K，其余的字符重复相同的操作。
现在牛牛尝试解密文件，但不知道所用密钥的值。鉴于文件的前8个字符总是以CHICKENS开头(因此该文件的key唯一)，现在请你编写一个程序,
将该字母文件解密为原始文件?
输入描述
在第一行中,给出了已执行XOR操作的字符串T
输出描述
输出执行XOR操作前的字符串S.

*/
func main() {
	// XOR运算 就是 异或运算！！
	var str string
	fmt.Scanln(&str)
	fmt.Println(myFunc(str))
}

func myFunc(strs string) (result string) {
	startStr := "CHICKENS"
	resBytes := make([]byte, len(strs))
	key := startStr[0] ^ strs[0]
	for i := 0; i < len(strs); i++ {
		resBytes[i] = strs[i] ^ key
	}
	return string(resBytes)
}
