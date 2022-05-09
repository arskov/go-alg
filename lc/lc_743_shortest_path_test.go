package lc

import (
	"container/heap"
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode743_0(t *testing.T) {
	in := [][]int{
		{2,1,1},
		{2,3,1},
		{3,4,1},
	}
	res := networkDelayTime(in, 4, 2)
	assert.Equal(t, 2, res)
}

func TestLeetcode743_1(t *testing.T) {
	in := [][]int{
		{1,2,1},
	}
	res := networkDelayTime(in, 2, 1)
	assert.Equal(t, 1, res)
}

func TestLeetcode743_2(t *testing.T) {
	in := [][]int{
		{1,2,1},
	}
	res := networkDelayTime(in, 2, 2)
	assert.Equal(t, -1, res)
}

func networkDelayTime(times [][]int, n int, k int) int {
	adjList := make([][]priorityqueue.Tuple2, n + 1)
    for _, edge := range times {
		adjList[edge[0]] = append(adjList[edge[0]], priorityqueue.Tuple2{Key: edge[2], First: edge[1]})
	}
	dist := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	pq := &priorityqueue.PQ{}
	//visited := make([]int, n + 1)
	heap.Push(pq, priorityqueue.Tuple2{Key: 0, First: k})
	for len(*pq) > 0 {
		top := heap.Pop(pq).(priorityqueue.Tuple2)
		for _, e := range adjList[top.First] {
			if e.Key + dist[top.First] < dist[e.First] {
				dist[e.First] = e.Key + dist[top.First]
				heap.Push(pq, e)
			}
		}
	}
	max := 0
	for i := 1; i <= n; i++ {
		if dist[i] > max {
			max = dist[i]
		}
	}
	if max == math.MaxInt32 {
		return -1
	} else {
		return max
	}
}