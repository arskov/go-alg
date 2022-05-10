package lc

import (
	"container/heap"
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode787_0(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	res := findCheapestPrice_1(4, in, 0, 3, 1)
	assert.Equal(t, 700, res)
}

func TestLeetcode787_1(t *testing.T) {
	in := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{3, 7, 1},
		{0, 4, 3},
		{4, 5, 3},
		{5, 7, 3},
		{0, 6, 5},
		{6, 7, 100},
		{7, 8, 1},
	}
	res := findCheapestPrice_1(9, in, 0, 8, 3)
	assert.Equal(t, 10, res)
}

func findCheapestPrice_1(n int, flights [][]int, src int, dst int, k int) int {
	adjList := make([][][]int, n)
	for _, f := range flights {
		adjList[f[0]] = append(adjList[f[0]], []int{f[1], f[2]})
	}
	dist := make([]int, n)
	stops := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		stops[i] = math.MaxInt32
	}
	dist[src] = 0
	stops[src] = 0
	pq := &priorityqueue.PQ{}
	heap.Push(pq, priorityqueue.Tuple3{Key:0, First: src, Second: 0})
	for pq.Len() > 0 {
		head := heap.Pop(pq).(priorityqueue.Tuple3)
		if head.First == dst {
			return head.Key
		}
		if head.Second == k+1 {
			continue
		}
		for _, adj := range adjList[head.First] {
			dU := head.Key
			dV := dist[adj[0]]
			wUV := adj[1]

			if dU+wUV < dV {
				heap.Push(pq, priorityqueue.Tuple3{Key: dU + wUV, First: adj[0], Second: head.Second + 1})
				dist[adj[0]] = dU + wUV
				stops[adj[0]] = head.Second
			} else if head.Second < stops[adj[0]] {
				heap.Push(pq, priorityqueue.Tuple3{Key: dU + wUV, First: adj[0], Second: head.Second + 1})
			}
		}
	}
	if dist[dst] == math.MaxInt32 {
		return -1
	} else {
		return dist[dst]
	}
}