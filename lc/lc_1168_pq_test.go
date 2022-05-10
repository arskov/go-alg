package lc

import (
	"container/heap"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mheap"
	"github.com/stretchr/testify/assert"
)

func TestLeetCode1168_0(t *testing.T) {
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
	assert.Equal(t, 3, minCostToSupplyWater_pq(len(wells), wells, pipes))
}

func minCostToSupplyWater_pq(n int, wells []int, pipes [][]int) int {
	graph := make(map[int][]mheap.Tuple2)
	pq := &mheap.PQ{}
	for i := 0; i < len(wells); i++ {
		pair := mheap.Tuple2{First: wells[i], Second: i + 1}
		graph[0] = append(graph[0], pair)
		*pq = append(*pq, pair)
	}
	for i := 0; i < len(pipes); i++ {
		h1 := pipes[i][0]
		h2 := pipes[i][1]
		newCost := pipes[i][2]
		graph[h1] = append(graph[h1], mheap.Tuple2{First: newCost, Second: h2})
		graph[h2] = append(graph[h2], mheap.Tuple2{First: newCost, Second: h1})
	}
	visited := make(map[int]struct{})
	heap.Init(pq)
	totalCost := 0

	for pq.Len() > 0 {
		pair := heap.Pop(pq).(mheap.Tuple2)
		if _, ok := visited[pair.Second]; !ok {
			visited[pair.Second] = struct{}{}
			totalCost += pair.First
			for _, otherPair := range graph[pair.Second] {
				if _, ok := visited[otherPair.Second]; !ok {
					heap.Push(pq, otherPair)
				}
			}
		}
	}

	return totalCost
}
