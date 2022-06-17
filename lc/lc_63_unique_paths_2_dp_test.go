package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode63_1_0(t *testing.T) {
	obstacleGrid := [][]int{
		{0},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	assert.Equal(t, 1, res)
}

func TestLeetcode63_1_1(t *testing.T) {
	obstacleGrid := [][]int{
		{1},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	assert.Equal(t, 0, res)
}

func TestLeetcode63_1_3(t *testing.T) {
	obstacleGrid := [][]int{
		{1},
		{0},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	assert.Equal(t, 0, res)
}

func TestLeetcode63_1_4(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 1},
		{0, 0},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	assert.Equal(t, 1, res)
}

func TestLeetcode63_1_5(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles(obstacleGrid)
	assert.Equal(t, 2, res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	grid[0][0] = obstacleGrid[0][0] ^ 1
	for i := 1; i < m; i++ {
		if grid[i-1][0] == 0 || obstacleGrid[i][0] == 1 {
			break
		}
		grid[i][0] = 1
	}
	for i := 1; i < n; i++ {
		if grid[0][i-1] == 0 || obstacleGrid[0][i] == 1 {
			break
		}
		grid[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				grid[i][j] = 0
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}
	return grid[m-1][n-1]
}
