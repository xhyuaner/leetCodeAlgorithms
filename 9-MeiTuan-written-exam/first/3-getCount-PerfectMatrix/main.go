package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	sc := bufio.NewScanner(os.Stdin)
	doubleArr := make([][]byte, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		for _, v := range strs {
			doubleArr[i] = append(doubleArr[i], []byte(v)...)
		}
	}
	fmt.Println(0)
	numsPerfect := 0
	for i := 1; i < n; i++ {
		numsPerfect = 0
		for x := 0; x < len(doubleArr); x++ {
			for y := 0; y < len(doubleArr[0]); y++ {
				if x+i < len(doubleArr) && y+i < len(doubleArr[0]) {
					if isPerfectMatrix(doubleArr, x, y, i) {
						numsPerfect++
					}
				}
			}
		}
		fmt.Println(numsPerfect)
	}

}

func isPerfectMatrix(arr [][]byte, row, col, n int) bool {
	count0, count1 := 0, 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if arr[row+i][col+j] == '0' {
				count0++
			} else {
				count1++
			}
		}
	}
	return count0 == count1
}
