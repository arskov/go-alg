package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode68_1_0(t *testing.T) {
	res := uniquePaths(3, 7)
	assert.Equal(t, 28, res)
}

func TestLeetcode68_1_1(t *testing.T) {
	res := uniquePaths(3, 2)
	assert.Equal(t, 3, res)
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		grid[i][0] = 1
	}
	for i := 0; i < n; i++ {
		grid[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}
