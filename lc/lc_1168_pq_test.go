package lc

import (
	"container/heap"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestLeetCode1168_0(t *testing.T) {
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
	assert.Equal(t, 3, minCostToSupplyWater_pq(len(wells), wells, pipes))
}

func minCostToSupplyWater_pq(n int, wells []int, pipes [][]int) int {
	graph := make(map[int][]priorityqueue.Tuple2)
	pq := &priorityqueue.PQ{}
	for i := 0; i < len(wells); i++ {
		pair := priorityqueue.Tuple2{Key: wells[i], First: i + 1}
		graph[0] = append(graph[0], pair)
		*pq = append(*pq, pair)
	}
	for i := 0; i < len(pipes); i++ {
		h1 := pipes[i][0]
		h2 := pipes[i][1]
		newCost := pipes[i][2]
		graph[h1] = append(graph[h1], priorityqueue.Tuple2{Key: newCost, First: h2})
		graph[h2] = append(graph[h2], priorityqueue.Tuple2{Key: newCost, First: h1})
	}
	visited := make(map[int]struct{})
	heap.Init(pq)
	totalCost := 0

	for pq.Len() > 0 {
		pair := heap.Pop(pq).(priorityqueue.Tuple2)
		if _, ok := visited[pair.First]; !ok {
			visited[pair.First] = struct{}{}
			totalCost += pair.Key
			for _, otherPair := range graph[pair.First] {
				if _, ok := visited[otherPair.First]; !ok {
					heap.Push(pq, otherPair)
				}
			}
		}
	}

	return totalCost
}
