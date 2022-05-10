package mheap

type PQ []Orderable

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].OrderKey() < pq[j].OrderKey()
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(v interface{}) {
	*pq = append(*pq, v.(Orderable))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	r := old[n-1]
	*pq = old[0 : n-1]
	return r
}
