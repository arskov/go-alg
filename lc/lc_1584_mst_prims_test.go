package lc

import (
	"container/heap"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/manhattan"
	"github.com/ArseniKavalchuk/dsa-go/pkg/mheap"
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
	pq := &mheap.PQ{}
	heap.Push(pq, mheap.Tuple2{First: 0, Second: 0})
	totalCost := 0
	count := 0
	for count < n {
		top := heap.Pop(pq).(mheap.Tuple2)
		if visited[top.Second] == 0 {
			totalCost += top.First
			visited[top.Second] = 1
			count++
			for i := 0; i < n; i++ {
				if i != top.Second && visited[i] == 0 {
					heap.Push(pq, mheap.Tuple2{First: manhattan.Dist(points[top.Second], points[i]), Second: i})
				}
			}
		}
	}
	return totalCost
}
