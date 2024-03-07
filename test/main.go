package main

import "fmt"

//func main() {
//	//tmp := []int{3, 1, 4}
//	//res := make([][]int, 0)
//	//
//	//res = append(res, tmp)
//	//res = append(res, []int{5})
//	//
//	//tmp[2] = 0
//	//fmt.Println(res) // 1
//	//
//	//tmp2 := []int{3, 1, 4}
//	//res = append(res, tmp2)
//	//tmp2[1] = 5
//	//
//	//fmt.Println(res) // 2
//}

// -----------------------------------------
//func main() {
//	arr := make([]int, 3, 4)             //创建一个长度为3，容量为4的切片
//	fmt.Println(arr, len(arr), cap(arr)) //[0 0 0] 3 4
//	// -----
//	addNum(arr)
//	// -----
//	fmt.Println(arr, len(arr), cap(arr)) //[0 0 0] 3 4
//}
//
//func addNum(sli []int) {
//	sli = append(sli, 4)
//	fmt.Println(sli, len(sli), cap(sli)) //[0 0 0 4] 4 4
//}
// ----------------------------------------------------------------
func main() {
	arr := make([]int, 3, 3) //创建一个长度为3，容量为4的切片
	//fmt.Printf("%p\n", arr)  //0xc000012200
	//fmt.Println(arr)         //[0 0 0]
	// -------
	addNum(arr)
	// -------
	//fmt.Printf("%p\n", arr) //0xc000012200
	//fmt.Println(arr)        //[0 0 0]
}

func addNum(sli []int) {
	fmt.Printf("%p\n", &sli)
	sli = append(sli, 4)
	fmt.Printf("%p\n", &sli) //0xc000012200
	//fmt.Println(sli)        //[0 0 0 4]
}
