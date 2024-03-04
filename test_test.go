package main

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {

	ch := make(chan int)

	close(ch)
	close(ch) // 关闭已经关闭的通道会引发 panic 错误

	value, ok := <-ch
	fmt.Println(value, ok)
	/**
	输出：
	panic: close of closed channel [recovered]
	panic: close of closed channel
	*/
}

//func TestF1(t *testing.T) {
//	s := map[string]int{}
//	s["a"]++
//	s["b"]++
//	fmt.Println(s)
//	fmt.Println("------------------")
//	s["a"]--
//	s["x"]--
//	// delete(s, "a")
//	fmt.Println(s)
//	/**	结果：
//	  map[a:1 b:1]
//	  ------------------
//	  map[a:0 b:1 x:-1]
//	// map[b:1 x:-1]
//	*/
//}

//func F1() int {
//	str := "abcde"
//	return str[2]
//}
