package main

import (
	"time"
)

// ----------------------test: for+select+time.After() 导致内存泄露-------------
func main() {
	ch := make(chan int, 10)
	go func() {
		i := 1
		for {
			i++
			ch <- i
		}
	}()

	for {
		select {
		case <-ch:
			// 当前case很快执行完毕，进入下次for循环时会再次创建一个新的time.After()，最后不断创建了大量的After()导致内存泄露
			// println(x)
			continue
		case <-time.After(3 * time.Minute): // 推荐使用time.NewTimer替代
			println(time.Now().Unix())
		}
	}
}

// // ----------------------test: sync.Map遍历方法-------------
// var m sync.Map

// func main() {
// 	// 写入
// 	data := []string{"煎鱼", "咸鱼", "烤鱼", "蒸鱼"}
// 	for i := 0; i < 4; i++ {
// 		go func(i int) {
// 			m.Store(i, data[i])
// 		}(i)
// 	}
// 	time.Sleep(time.Second)
// 	// 遍历
// 	m.Range(func(key, value interface{}) bool {
// 		fmt.Printf("Range: %v, %v\n", key, value)
// 		return true // Range: 3, 蒸鱼	Range: 0, 煎鱼	Range: 1, 咸鱼	Range: 2, 烤鱼
// 		// return false // Range: 1, 咸鱼
// 	})
// }

// // ----------------------test: 空结构体与非空结构体地址-------------
// type People struct {
// 	Name string
// }

// func main() {
// 	a := &People{Name: "liu"}
// 	b := &People{Name: "liu"}
// 	fmt.Printf("%p\n", a)
// 	fmt.Printf("%p\n", b)
// 	fmt.Println(a == b)
// }

// ----------------------test: 无缓冲通道---阻塞读写-------------
// func main() {
// 	ch := make(chan int)
// 	ch <- 2
// 	<-ch
// }
// --------------------------test00---------------------------
// func main() {
// 	ch := make(chan string, 5)
// 	ch <- "first"
// 	ch <- "second"
// 	close(ch)
// 	x, ok := <-ch
// 	if ok {
// 		fmt.Println("received1: ", x)
// 	}

// 	x, ok = <-ch
// 	if ok {
// 		fmt.Println("received2: ", x)
// 	} else {
// 		fmt.Println("channel closed, data invalid.", x)
// 	}

// 	x, ok = <-ch
// 	if !ok {
// 		fmt.Println("channel closed, data invalid.", x)
// 	} else {
// 		fmt.Println("received3: ", x)
// 	}
// }
