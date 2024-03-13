package main

import (
	"fmt"
)

func main() {
	// 原始的slice
	originalSlice := []int{1, 2, 3, 4, 5}

	// 新的slice，用来存放指针
	pointerSlice := make([]*int, len(originalSlice))

	// 遍历原始的slice，将每个元素的地址赋给新的slice
	for i := range originalSlice {
		pointerSlice[i] = &originalSlice[i]
	}

	// 打印新的slice中的指针值
	for _, ptr := range pointerSlice {
		fmt.Println(ptr) // 输出的是每个指针指向的值
	}
}
