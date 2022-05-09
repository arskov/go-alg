package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode994_1(t *testing.T) {
	in := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	res := orangesRotting(in)
	assert.Equal(t, 4, res)
}

func TestLeetcode994_2(t *testing.T) {
	in := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	res := orangesRotting(in)
	assert.Equal(t, -1, res)
}

func TestLeetcode994_3(t *testing.T) {
	in := [][]int{
		{0, 2},
	}
	res := orangesRotting(in)
	assert.Equal(t, 0, res)
}

var dir4 = [][]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	queue := make([][]int, 0)
	h := len(grid)
	w := len(grid[0])
	freshCount := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch grid[i][j] {
			case 1:
				freshCount += 1
			case 2:
				queue = append(queue, []int{i, j})
			}
		}
	}
	count := -1
	for len(queue) > 0 {
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			head := queue[0]
			queue = queue[1:]
			for _, d := range dir4 {
				x := head[0] + d[0]
				y := head[1] + d[1]
				if x >= 0 && x < h && y >= 0 && y < w && grid[x][y] == 1 {
					grid[x][y] = 2
					freshCount -= 1
					queue = append(queue, []int{x, y})
				}
			}
		}
		count += 1
	}
	if freshCount > 0 {
		return -1
	} else {
		return count
	}
}
