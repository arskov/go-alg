package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1092_1(t *testing.T) {
	in := [][]int{
		{0, 1},
		{1, 0},
	}
	res := shortestPathBinaryMatrix(in)
	assert.Equal(t, 2, res)
}

func TestLeetcode1092_2(t *testing.T) {
	in := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	res := shortestPathBinaryMatrix(in)
	assert.Equal(t, 4, res)
}

func TestLeetcode1092_3(t *testing.T) {
	in := [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	res := shortestPathBinaryMatrix(in)
	assert.Equal(t, -1, res)
}

var dir8 = [][]int{
	{-1, -1},
	{1, -1},
	{1, 1},
	{-1, 1},
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || grid[0][0] == 1 {
		return -1
	}
	last := len(grid) - 1
	pathLen := 1
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	grid[0][0] = 1
	for len(queue) > 0 {
		levLen := len(queue)
		for i := 0; i < levLen; i++ {
			head := queue[0]
			queue = queue[1:]
			if head[0] == last && head[1] == last {
				return pathLen
			}
			for _, d := range dir8 {
				x := head[0] + d[0]
				y := head[1] + d[1]
				if x >= 0 && x <= last && y >= 0 && y <= last && grid[x][y] == 0 {
					grid[x][y] = 1
					queue = append(queue, []int{x, y})
				}
			}
		}
		pathLen += 1
	}
	return -1
}
