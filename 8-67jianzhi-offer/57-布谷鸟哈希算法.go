package jianzhi

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TableSize = 11  // 哈希表的大小
	MaxKicks  = 500 // 重新放置的最大次数
)

// CuckooHashTable 布谷鸟哈希表结构
type CuckooHashTable struct {
	table1 []int
	table2 []int
}

// NewCuckooHashTable 创建新的布谷鸟哈希表
func NewCuckooHashTable() *CuckooHashTable {
	return &CuckooHashTable{
		table1: make([]int, TableSize),
		table2: make([]int, TableSize),
	}
}

// hash1 第一个哈希函数
func hash1(key int) int {
	return key % TableSize
}

// hash2 第二个哈希函数
func hash2(key int) int {
	return (key / TableSize) % TableSize
}

// Insert 插入元素到哈希表中
func (h *CuckooHashTable) Insert(key int) bool {
	for i := 0; i < MaxKicks; i++ {
		if h.table1[hash1(key)] == 0 {
			h.table1[hash1(key)] = key
			return true
		} else if h.table2[hash2(key)] == 0 {
			h.table2[hash2(key)] = key
			return true
		} else {
			// 选择随机表进行替换
			if rand.Intn(2) == 0 {
				key, h.table1[hash1(key)] = h.table1[hash1(key)], key
			} else {
				key, h.table2[hash2(key)] = h.table2[hash2(key)], key
			}
		}
	}
	return false
}

// Search 搜索元素
func (h *CuckooHashTable) Search(key int) bool {
	return h.table1[hash1(key)] == key || h.table2[hash2(key)] == key
}

// Delete 删除元素
func (h *CuckooHashTable) Delete(key int) {
	if h.table1[hash1(key)] == key {
		h.table1[hash1(key)] = 0
	} else if h.table2[hash2(key)] == key {
		h.table2[hash2(key)] = 0
	}
}

// Display 打印哈希表
func (h *CuckooHashTable) Display() {
	fmt.Println("Table 1:")
	for i, v := range h.table1 {
		fmt.Printf("Index %d: %d\n", i, v)
	}
	fmt.Println("Table 2:")
	for i, v := range h.table2 {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	h := NewCuckooHashTable()

	// 插入元素
	h.Insert(10)
	h.Insert(22)
	h.Insert(31)
	h.Insert(4)
	h.Insert(15)
	h.Insert(28)
	h.Insert(17)
	h.Insert(88)
	h.Insert(59)

	h.Display()

	// 搜索元素
	fmt.Println("Search 31:", h.Search(31))
	fmt.Println("Search 99:", h.Search(99))

	// 删除元素
	h.Delete(31)
	fmt.Println("After deleting 31:")
	h.Display()
}
