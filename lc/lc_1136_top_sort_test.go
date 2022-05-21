package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1136_1_0(t *testing.T) {
	n := 3
	relations := [][]int{
		{1, 3},
		{2, 3},
	}
	res := minimumSemesters(n, relations)
	assert.Equal(t, 2, res)
}

func TestLeetcode1136_1_1(t *testing.T) {
	n := 3
	relations := [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	res := minimumSemesters(n, relations)
	assert.Equal(t, -1, res)
}

func minimumSemesters(n int, relations [][]int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for _, edge := range relations {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}
	queue := make([]int, 0)
	for i := 1; i < n+1; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	nodeCount := 0
	semesterCount := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		semesterCount++
		for i := 0; i < levelSize; i++ {
			nodeCount++
			top := queue[0]
			queue = queue[1:]
			for _, adj := range graph[top] {
				inDegree[adj]--
				if inDegree[adj] == 0 {
					queue = append(queue, adj)
				}
			}
		}
	}
	if nodeCount == n {
		return semesterCount
	} else {
		return -1
	}
}
