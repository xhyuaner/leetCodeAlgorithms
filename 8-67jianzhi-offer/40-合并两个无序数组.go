package jianzhi

import "fmt"

/**
 * mergeArray
 *  @Description: 合并两个无序数组，递增返回k个最大值 --- 归并排序
 *  @param a
 *  @param b
 *  @param k
 *  @return ans
 */
func mergeArray(a, b []int, k int) (ans []int) {
	var sort func(arr []int) []int
	sort = func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}
		mid := len(arr) / 2
		left := sort(arr[:mid])
		right := sort(arr[mid:])
		return merge2(left, right)
	}
	le := sort(a)
	ri := sort(b)
	resArr := merge2(le, ri)
	for i := 0; i < k; i++ {
		ans = append(ans, resArr[i])
	}
	return ans
}

func merge2(left, right []int) (ans []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			ans = append(ans, left[l])
			l++
		} else {
			ans = append(ans, right[r])
			r++
		}
	}
	ans = append(ans, left[l:]...)
	ans = append(ans, right[r:]...)
	return ans
}
func main() {
	arr1 := []int{38, 27, 43}
	arr2 := []int{3, 9, 82, 10}
	//fmt.Println("Original array:", arr)
	res := mergeArray(arr1, arr2, 3)
	fmt.Println("Sorted array:", res)
}
