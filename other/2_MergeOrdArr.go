package other

/*
【两个有序数组合并】
问题描述:
给定两个有序整数数组 A 和 B，将B合并到A中，使得 A 成为一个有序数组。
说明:
1. 初始化 A 和 B 的元素数量分别为 m 和 n。
2. A有足够的空间（空间大小大于或等于 m + n）来保存 B 中的元素。
3. 默认升序。

【例如】
输入:
数组A，以及数组A元素数量
数组B，以及数组B元素数量
A = [1,6,7,0,0,0], m = 3
B = [2,4,6], n = 3
输出：
合并后的数组A
A = [1,2,4,6,6,7];
*/
import "fmt"

func merge(arr1, arr2 []int, m, n int) {
	// 两个数组从后往前遍历合并
	i, j, idx := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if arr1[i] > arr2[j] {
			arr1[idx] = arr1[i]
			idx--
			i--
		} else {
			arr1[idx] = arr2[j]
			idx--
			j--
		}
	}
	// arr2有剩余（由于是有序的，所以arr1有剩余不用处理，剩下各元素直接在原位置即可）
	for ; j >= 0; j-- {
		arr1[idx] = arr2[j]
		idx--
	}
	fmt.Println("合并结果为：", arr1)
}

//func main() {
//	var m,n int
//	fmt.Println("请分别输入两个有序数组大小m和n：")
//	fmt.Scanf("%d %d\n",&m,&n)
//	// make()函数初始化int数组各元素默认为0
//	arr1:=make([]int,m+n,m+n)
//	arr2:=make([]int,n,n)
//	fmt.Println("请输入第一个有序数组各元素：")
//	for i:=0;i<m;i++{
//		fmt.Scan(&arr1[i])
//	}
//	fmt.Println("请输入第二个有序数组各元素：")
//	for i:=0;i<n;i++{
//		fmt.Scan(&arr2[i])
//	}
//	// 调用merge()函数
//	merge(arr1, arr2, m, n)
//}
