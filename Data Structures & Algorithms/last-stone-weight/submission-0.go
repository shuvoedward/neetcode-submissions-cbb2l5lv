type MaxHeap []int

func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Swap(i, j int){h[i], h[j] = h[j], h[i]}
func (h MaxHeap) Less(i, j int)bool{return h[i] > h[j]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(int))}
func (h *MaxHeap) Pop()interface{}{
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}


func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, s := range stones{
		heap.Push(h, s)
	}

	for h.Len() > 1{
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)
		if y != x{
			heap.Push(h, y-x)
		}
	}

	if h.Len() == 0{
		return 0
	}
	return (*h)[0]
}
