package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode265_1_0(t *testing.T) {
	costs := [][]int{
		{1, 5, 3},
		{2, 9, 4},
	}
	res := minCostII_265(costs)
	assert.Equal(t, 5, res)
}

func TestLeetcode265_1_1(t *testing.T) {
	costs := [][]int{
		{1, 3},
		{2, 4},
	}
	res := minCostII_265(costs)
	assert.Equal(t, 5, res)
}

func minCostII_265(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])
	memo := make([][]int, k)
	for i := 0; i < k; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(step, color int) int
	dp = func(step, color int) int {
		if step == n {
			return 0
		}
		if memo[color][step] == -1 {
			var best int = math.MaxInt32
			for i := 1; i < k; i++ {
				tmp := dp(step+1, (color+i)%k)
				if tmp < best {
					best = tmp
				}
			}
			memo[color][step] = costs[step][color] + best
		}
		return memo[color][step]
	}
	var best int = math.MaxInt32
	for i := 0; i < k; i++ {
		tmp := dp(0, i)
		if tmp < best {
			best = tmp
		}
	}
	return best
}
