package jianzhi

//func main() {
//	//arr := []int{7, 6, 1, 2, 3, 4, 5, 9}
//	arr := []int{2, 1, 3, 5, 7}
//	fmt.Println(arr)
//	reOrderArray(arr)
//	fmt.Println(arr)
//}

func reOrderArray(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		for l < r && arr[l]%2 == 1 {
			l++
		}
		for l < r && arr[r]%2 == 0 {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
}
