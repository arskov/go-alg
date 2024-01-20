package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode256_1_0(t *testing.T) {
	costs := [][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}
	res := minCostI_256(costs)
	assert.Equal(t, 10, res)
}

func TestLeetcode256_1_1(t *testing.T) {
	costs := [][]int{
		{7, 6, 2},
	}
	res := minCostI_256(costs)
	assert.Equal(t, 2, res)
}

func minCostI_256(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	memo := make([][]int, 3)
	memo[0] = make([]int, n)
	memo[1] = make([]int, n)
	memo[2] = make([]int, n)
	for i := 0; i < n; i++ {
		memo[0][i] = -1
		memo[1][i] = -1
		memo[2][i] = -1
	}
	var dp func(step, nextColor int) int
	dp = func(step, nextColor int) int {
		if step == n {
			return 0
		}
		if memo[nextColor][step] == -1 {
			memo[nextColor][step] = costs[step][nextColor] +
				min(dp(step+1, (nextColor+1)%3), dp(step+1, (nextColor+2)%3))
		}
		return memo[nextColor][step]
	}
	return min(dp(0, 0), min(dp(0, 1), dp(0, 2)))
}
