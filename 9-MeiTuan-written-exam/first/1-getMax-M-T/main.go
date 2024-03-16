package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, count int
	fmt.Scanln(&n, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	for i := 0; i < len(str); i++ {
		if str[i] == 'M' || str[i] == 'T' {
			continue
		}
		count++

	}
	if k <= count {
		fmt.Println(n - count + k)
	} else {
		fmt.Println(n)
	}

}
