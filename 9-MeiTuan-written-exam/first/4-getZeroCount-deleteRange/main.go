package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func myFun(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			var newArr []int
			result := 1
			newArr = append(newArr, arr[:i]...)
			newArr = append(newArr, arr[j+1:]...)
			for _, v := range newArr {
				result *= v
			}
			resultStr := strconv.Itoa(result)
			for i := len(resultStr) - 1; resultStr[i] == '0'; i-- {
				if i == len(resultStr)-k {
					count++
					break
				}
			}
		}
	}
	return count
}

func main() {
	var n, k int
	fmt.Scanln(&n, &k)

	arr := make([]int, n)
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 读取一行输入，用空格隔开
	input = scanner.Text()

	strArr := strings.Split(input, " ")
	for i, str := range strArr {
		arr[i], _ = strconv.Atoi(str)
	}

	//fmt.Println("输入的数组:", arr)
	//fmt.Println("输入的整数k:", k)
	fmt.Println(myFun(arr, k))
}
