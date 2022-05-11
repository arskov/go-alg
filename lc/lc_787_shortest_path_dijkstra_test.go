package lc

import (
	"container/heap"
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mheap"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode787_1_0(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	res := findCheapestPrice_dijkstra(4, in, 0, 3, 1)
	assert.Equal(t, 700, res)
}

func TestLeetcode787_1_1(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_dijkstra(3, in, 0, 2, 1)
	assert.Equal(t, 200, res)
}

func TestLeetcode787_1_2(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_dijkstra(3, in, 0, 2, 0)
	assert.Equal(t, 500, res)
}

func TestLeetcode787_1_3(t *testing.T) {
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
	res := findCheapestPrice_dijkstra(9, in, 0, 8, 3)
	assert.Equal(t, 10, res)
}

// Dijkstra version of the algorithm doesn't work for negative edges
//
// func TestLeetcode787_1_4(t *testing.T) {
// 	in := [][]int{
// 		{1, 2, 100},
// 		{3, 1, -150},
// 		{0, 1, 100},
// 		{2, 3, 100},
// 		{0, 3, 200},
// 		{0, 2, 500},
// 	}
// 	res := findCheapestPrice_dijkstra(4, in, 0, 1, 3)
// 	assert.Equal(t, 50, res)
// }

func findCheapestPrice_dijkstra(n int, flights [][]int, src int, dst int, k int) int {
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
	pq := &mheap.PQ{}
	heap.Push(pq, mheap.Tuple3{First: 0, Second: src, Third: 0})
	for pq.Len() > 0 {
		head := heap.Pop(pq).(mheap.Tuple3)
		if head.Second == dst {
			return head.First
		}
		if head.Third == k+1 {
			continue
		}
		for _, adj := range adjList[head.Second] {
			dU := head.First
			dV := dist[adj[0]]
			wUV := adj[1]
			if dU+wUV < dV {
				heap.Push(pq, mheap.Tuple3{First: dU + wUV, Second: adj[0], Third: head.Third + 1})
				dist[adj[0]] = dU + wUV
				stops[adj[0]] = head.Third
			} else if head.Third < stops[adj[0]] {
				heap.Push(pq, mheap.Tuple3{First: dU + wUV, Second: adj[0], Third: head.Third + 1})
			}
		}
	}
	if dist[dst] == math.MaxInt32 {
		return -1
	} else {
		return dist[dst]
	}
}
