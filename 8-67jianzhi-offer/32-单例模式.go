package jianzhi

import "sync"

// 饿汉式单例模式
type singleton struct{}

var instance = &singleton{}

func GetInstance() *singleton {
	return instance
}

// 懒汉式单例模式，注意线程安全问题
type singleton2 struct{}

var (
	instance2 *singleton2
	mutex     sync.Mutex
)

func GetInstance2() *singleton2 {
	// 第一次判断nil：如果不是nil，就直接返回实例，避免每次都要拿锁，影响性能
	if instance2 == nil {
		// 如果实例是nil，就锁住创建实例
		mutex.Lock()
		defer mutex.Unlock()
		// 第二次判断nil：为了确保即便是有多个协程绕过了第一次nil检查， 也只能有一个可以创建单例实例。
		// 否则， 多个协程都会创建自己的实例并返回
		if instance2 == nil {
			instance2 = &singleton2{}
			return instance2
		} else {
			return instance2
		}
	} else {
		return instance2
	}
}
