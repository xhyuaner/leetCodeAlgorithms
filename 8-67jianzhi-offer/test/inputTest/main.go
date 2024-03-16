package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum, num int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sum = 0
		strs := strings.Split(scanner.Text(), " ")
		for _, v := range strs {
			num, _ = strconv.Atoi(v)
			sum += num
		}
		fmt.Println(sum)
	}
}
