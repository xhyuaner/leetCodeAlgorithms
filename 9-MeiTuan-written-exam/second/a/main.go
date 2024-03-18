package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	sum := 0
	price := 0
	for _, v := range strs {
		price, _ = strconv.Atoi(v)
		sum += price
	}
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(sum - x - y)
}
