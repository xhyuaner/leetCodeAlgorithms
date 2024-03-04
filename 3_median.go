package main

// --------------------解法一：二分查找两有序数组中位数---------------------------
// func findMedianSortArrays(nums1, nums2 []int) float64 {
// 	totalLength := len(nums1) + len(nums2)
// 	if totalLength%2 == 1 {
// 		midIndex := totalLength / 2
// 		return float64(getKthElement(nums1, nums2, midIndex+1))
// 	} else {
// 		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
// 		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
// 	}
// }

// func getKthElement(nums1, nums2 []int, k int) int {
// 	index1, index2 := 0, 0
// 	for {
// 		if index1 == len(nums1) {
// 			return nums2[index2+k-1]
// 		}
// 		if index2 == len(nums2) {
// 			return nums1[index1+k-1]
// 		}
// 		if k == 1 {
// 			return min(nums1[index1], nums2[index2])
// 		}
// 		harf := k / 2
// 		newIndex1 := min(index1+harf, len(nums1)) - 1
// 		newIndex2 := min(index2+harf, len(nums2)) - 1
// 		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
// 		if pivot1 <= pivot2 {
// 			k = k - (newIndex1 - index1 + 1)
// 			index1 = newIndex1 + 1
// 		} else {
// 			k = k - (newIndex2 - index2 + 1)
// 			index2 = newIndex2 + 1
// 		}
// 	}
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// --------------------解法二：划分数组 + 二分查找 两有序数组中位数，时：O(log(min(m,n)))，空：O(1)---------------------------
func findMedianSortArrays(A, B []int) float64 {
	m, n := len(A), len(B)
	if m > n {
		return findMedianSortArrays(B, A)
	}
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i != m && B[j-1] > A[i] { // i 需要增大
			iMin = i + 1
		} else if i != 0 && j != n && A[i-1] > B[j] { // i 需要减小
			iMax = i - 1
		} else { // 达到要求，并且将边界条件列出来单独考虑
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = max(A[i-1], B[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft) // 奇数不用考虑右半部分
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = min(A[i], B[j])
			}
			return (float64(maxLeft) + float64(minRight)) / 2.0
		}
	}
	return 0.0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func main() {
// 	nums1 := []int{1, 3, 4, 9}
// 	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	result := findMedianSortArrays(nums1, nums2)
// 	fmt.Println(result)
// }
