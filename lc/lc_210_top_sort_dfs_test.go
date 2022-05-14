package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode210_2_0(t *testing.T) {
	in := [][]int{
		{0, 1},
	}
	res := findOrder_2(2, in)
	assert.Equal(t, []int{1, 0}, res)
}

func TestLeetcode210_2_1(t *testing.T) {
	in := [][]int{
		{1, 0},
	}
	res := findOrder_2(2, in)
	assert.Equal(t, []int{0, 1}, res)
}

func TestLeetcode210_2_2(t *testing.T) {
	in := [][]int{}
	res := findOrder_2(1, in)
	assert.Equal(t, []int{0}, res)
}

func TestLeetcode210_2_3(t *testing.T) {
	in := [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}
	res := findOrder_2(3, in)
	assert.Equal(t, []int{}, res)
}

func findOrder_2(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	} else if numCourses == 1 {
		return []int{0}
	}
	graph := make([][]int, numCourses)
	for _, e := range prerequisites {
		s := e[1]
		d := e[0]
		graph[s] = append(graph[s], d)
	}
	visited := make([]int, numCourses)
	order := make([]int, numCourses)
	j := numCourses - 1
	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = 1
		for _, next := range graph[node] {
			if visited[next] == 1 {
				return true
			}
			if visited[next] == 0 {
				res := dfs(next)
				if res {
					return true
				}
			}
		}
		order[j] = node
		j -= 1
		visited[node] = 2
		return false
	}
	cycle := false
	for i, v := range visited {
		if v == 0 {
			cycle = cycle || dfs(i)
			if cycle {
				break
			}
		}
	}
	if cycle {
		return []int{}
	} else {
		return order
	}
}
