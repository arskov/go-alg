package lc

import (
	"container/heap"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/manhattan"
	"github.com/ArseniKavalchuk/dsa-go/pkg/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode1584_2(t *testing.T) {
	in := [][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}
	res := minCostConnectPoints_2(in)
	assert.Equal(t, 20, res)
}

func minCostConnectPoints_2(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	n := len(points)
	visited := make([]int, n)
	pq := &priorityqueue.PQ{}
	heap.Push(pq, priorityqueue.Tuple2{Key: 0, First: 0})
	totalCost := 0
	count := 0
	for count < n {
		top := heap.Pop(pq).(priorityqueue.Tuple2)
		if visited[top.First] == 0 {
			totalCost += top.Key
			visited[top.First] = 1
			count++
			for i := 0; i < n; i++ {
				if i != top.First && visited[i] == 0 {
					heap.Push(pq, priorityqueue.Tuple2{Key: manhattan.Dist(points[top.First], points[i]), First: i})
				}
			}
		}
	}
	return totalCost
}
