package test

import (
	jz "LeetCodeAlgorithms/8-67jianzhi-offer"
	"fmt"
	"testing"
)

func TestReplaceWords(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	t.Log(jz.ReplaceWords(dictionary, sentence))
}

type OverviewResp struct {
	FileCount   int64 `json:"file_count"`
	MemberCount int64 `json:"member_count"`
}

func TestTrap(t *testing.T) {

	x := OverviewResp{
		FileCount:   66,
		MemberCount: 55,
	}
	y := OverviewResp{
		FileCount:   0,
		MemberCount: 0,
	}
	z := OverviewResp{}
	fmt.Println(x == y)
	fmt.Println(x == z)
	fmt.Println(y == z)

	//height := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//height := []int{2, 3, 5, 7, 9, 18, 101}
	//g := height[:0]
	//t.Log(jz.LengthOfLIS(height))
	//slice := make([]int, 5)
	////slice := []int{11, 22}
	//fmt.Println(slice, len(slice), cap(slice)) // 输出: [10]
	//push(slice, 77)
	//fmt.Println(slice, len(slice), cap(slice)) // 这行代码不会被执行
}

func push(slice []int, item int) {
	slice[3] = item
	//slice = append(slice, item)
	fmt.Println(slice, len(slice), cap(slice)) // 这行代码不会被执行
}

func TestTopK(t *testing.T) {
	//intervals := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	////nums := []int{1, 1, 1, 2, 2, 3}
	//t.Log(jz.MaxSubArray(intervals))
	arr1 := []int{38, 27, 43}
	arr2 := []int{3, 9, 82, 10}
	//fmt.Println("Original array:", arr)
	res := jz.MergeArray(arr1, arr2, 6)
	fmt.Println("Sorted array:", res)
}
