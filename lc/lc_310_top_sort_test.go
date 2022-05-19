package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode310_1_0(t *testing.T) {
	n := 1
	edges := [][]int{}
	res := findMinHeightTrees(n, edges)
	assert.Equal(t, []int{0}, res)
}

func TestLeetcode310_1_1(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
	}
	res := findMinHeightTrees(n, edges)
	assert.Equal(t, []int{1}, res)
}

func TestLeetcode310_1_2(t *testing.T) {
	n := 6
	edges := [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}
	res := findMinHeightTrees(n, edges)
	assert.Equal(t, []int{3, 4}, res)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degree := make([]int, n)

	for _, e := range edges {
		u := e[0]
		v := e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		degree[u] += 1
		degree[v] += 1
	}

	queue := make([]int, 0)
	for i, v := range degree {
		if v == 1 {
			queue = append(queue, i)
			//degree[i] -= 1
		}
	}
	remaining := n
	for remaining > 2 {
		queueSize := len(queue)
		remaining -= queueSize
		for i := 0; i < queueSize; i++ {
			top := queue[0]
			queue = queue[1:]
			degree[top] -= 1
			for _, adj := range graph[top] {
				if degree[adj] > 0 {
					degree[adj] -= 1
					if degree[adj] == 1 {
						queue = append(queue, adj)
					}
				}
			}
		}
	}
	return queue
}
