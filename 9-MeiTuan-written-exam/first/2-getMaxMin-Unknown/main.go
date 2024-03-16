package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q, sum, num, count int
	fmt.Scanln(&n, &q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	for _, v := range strs {
		num, _ = strconv.Atoi(v)
		if num == 0 {
			// 未知数数量
			count++
		} else {
			sum += num
		}
	}
	var l, r int
	for q > 0 {
		fmt.Scanln(&l, &r)
		fmt.Println(sum+count*l, sum+count*r)
	}
}
