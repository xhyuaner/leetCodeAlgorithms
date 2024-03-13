package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Pool struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

func NewPool(maxGoroutines int) *Pool {
	pool := &Pool{
		taskChan: make(chan Task),
	}
	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			defer pool.wg.Done()
			for task := range pool.taskChan {
				task()
			}
		}()
	}
	return pool
}

func (p *Pool) AddTask(task Task) {
	p.taskChan <- task
}

func (p *Pool) Shutdown() {
	close(p.taskChan)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5) // 创建一个拥有5个goroutine的池

	for i := 0; i < 10; i++ {
		count := i
		pool.AddTask(func() {
			fmt.Printf("Executing task %d\n", count)
			time.Sleep(1 * time.Second) // 模拟耗时操作
		})
	}

	pool.Shutdown() // 等待所有任务完成并关闭协程池
	fmt.Println("All tasks completed.")
}
