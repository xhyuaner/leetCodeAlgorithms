package jianzhi

import (
	"fmt"
)

func main() {
	arr := []int{7, 6, 1, 2, 3, 4, 5, 9}
	fmt.Println(arr)
	reOrderArray(arr)
	fmt.Println(arr)
}

func reOrderArray(arr []int) {
	for i := 0; i <= len(arr)/2; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j]%2 == 1 && arr[j-1]%2 == 0 {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
