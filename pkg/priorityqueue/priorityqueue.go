package priorityqueue

type PriorityQueue []Orderable

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].OrderKey() < pq[j].OrderKey()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(v interface{}) {
	*pq = append(*pq, v.(Orderable))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	r := old[n-1]
	*pq = old[0 : n-1]
	return r
}
