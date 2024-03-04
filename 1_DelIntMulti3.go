package main

/*
【整数列表求三的倍数】
问题描述:
1、给定一个从1到n的整数列表，从第一个数字开始计数，遇到3的倍数时，将该数从列表中删除，直至列表末尾。
2、在剩下的数字中，从第一个数字开始，继续之前的计数值，同样遇到3的倍数时，删除该数。
3、循环上面的步骤，直到列表中只剩下一个数字。
4、根据指定的数字n，来判断最后剩下的数字是哪个。

【例如】
输入:n = 5
第一步： 1, 2, (3), 4, 5
第二步： (1), 2, 4, (5)
第三步：(2), 4
输出：4
*/

import (
	"fmt"
)

// 整数列表求三的倍数函数
func getNum(n int) {
	// 初始化切片，通过make()创建的int切片各元素默认值为0
	array := make([]int, n, 10)
	index, count := 0, 0
	length := n
	fmt.Printf("输入的n是：%v,数组是：%v\n", n, array) // n=5,array=[0 0 0 0 0]
	// 当n为1时终止循环
	for n > 1 {
		// 通过123123123...循环计数来判断当前元素是否为3的倍数
		if array[index] == 0 {
			count++
			// 若是3的倍数，将元素值由默认的0置为1，并且将n减1
			if count == 3 {
				count = 0
				array[index] = 1
				n--
			}
		}
		index++
		// 当扫描完最后一个元素后，若n!=1则返回到初始元素继续重新扫描
		if index == length {
			index = 0
		}
	}
	// 遍历切片，输出元素值仍为0的那个元素即为最终结果
	for j := 0; j < length; j++ {
		if array[j] != 1 {
			fmt.Println(j + 1)
		}
	}
}

//func main() {
//	var n int
//	// 接收用户输入的n
//	_, err := fmt.Scanf("%v", &n)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("main函数的n是：", n)
//	// 调用函数
//	getNum(n)
//	return
//}
