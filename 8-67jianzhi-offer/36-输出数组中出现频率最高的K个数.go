package jianzhi

import "container/heap"

type Iheap [][2]int

func (h Iheap) Len() int { return len(h) }
func (h Iheap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h Iheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Iheap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *Iheap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/**
 * topKFrequent
 *  @Description: 输出前K个高频元素
 *  @param nums
 *  @param k
 *  @return []int
 */
func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	h := &Iheap{}
	heap.Init(h)
	for key, value := range mp {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := k; i > 0; i-- {
		res[i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
