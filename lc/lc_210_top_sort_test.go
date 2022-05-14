package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode210_1_0(t *testing.T) {
	in := [][]int{
		{0, 1},
	}
	res := findOrder(2, in)
	assert.Equal(t, []int{1, 0}, res)
}

func TestLeetcode210_1_1(t *testing.T) {
	in := [][]int{
		{1, 0},
	}
	res := findOrder(2, in)
	assert.Equal(t, []int{0, 1}, res)
}

func TestLeetcode210_1_2(t *testing.T) {
	in := [][]int{}
	res := findOrder(1, in)
	assert.Equal(t, []int{0}, res)
}

func TestLeetcode210_1_3(t *testing.T) {
	in := [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}
	res := findOrder(3, in)
	assert.Equal(t, []int{}, res)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	} else if numCourses == 1 {
		return []int{0}
	}
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, e := range prerequisites {
		s := e[1]
		d := e[0]
		graph[s] = append(graph[s], d)
		inDegree[d] += 1
	}
	visited := make([]int, numCourses)
	var detectCycle func(node int) bool
	detectCycle = func(node int) bool {
		for _, next := range graph[node] {
			if visited[next] == 1 {
				return true
			}
			if visited[next] == 0 {
				visited[next] = 1
				res := detectCycle(next)
				if res {
					return true
				}
				visited[next] = 2
			}
		}
		return false
	}
	cycle := false
	for i, v := range visited {
		if v == 0 {
			visited[i] = 1
			cycle = cycle || detectCycle(i)
			visited[i] = 2
		}
	}
	if cycle {
		return []int{}
	}
	queue := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	if len(queue) == 0 {
		return []int{}
	}
	order := make([]int, numCourses)
	i := 0
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		order[i] = head
		i++
		for _, adj := range graph[head] {
			inDegree[adj] -= 1
			if inDegree[adj] == 0 {
				queue = append(queue, adj)
			}
		}
	}
	return order
}
