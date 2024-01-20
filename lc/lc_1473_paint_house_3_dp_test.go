package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1473_1_0(t *testing.T) {
	houses := []int{0, 0, 0, 0, 0}
	cost := [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}
	m := 5
	n := 2
	target := 3
	res := minCostIII_1473(houses, cost, m, n, target)
	assert.Equal(t, 9, res)
}

func TestLeetcode1473_1_1(t *testing.T) {
	houses := []int{3, 1, 2, 3}
	cost := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	m := 4
	n := 3
	target := 3
	res := minCostIII_1473(houses, cost, m, n, target)
	assert.Equal(t, -1, res)
}

func TestLeetcode1473_1_2(t *testing.T) {
	houses := []int{0, 2, 1, 2, 0}
	cost := [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}
	m := 5
	n := 2
	target := 3
	res := minCostIII_1473(houses, cost, m, n, target)
	assert.Equal(t, 11, res)
}

func minCostIII_1473(houses []int, cost [][]int, m int, n int, target int) int {
	if len(houses) == 0 || n == 0 || target == 0 || len(cost) == 0 {
		return 0
	}
	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, target+1)
			for k := 0; k <= target; k++ {
				memo[i][j][k] = math.MinInt32
			}
		}
	}
	var dp func(step, prevColor, nCount int) int
	dp = func(step, prevColor, nCount int) int {
		if nCount > target {
			return math.MaxInt32
		}
		if step == m {
			if nCount != target {
				return math.MaxInt32
			}
			return 0
		}
		if prevColor == 0 || memo[step][prevColor-1][nCount] == math.MinInt32 {
			best := math.MaxInt32
			if houses[step] == 0 {
				for i := 1; i <= n; i++ {
					best = min(best, cost[step][i-1]+dp(step+1, i, nextNeighbour(nCount, prevColor, i)))
				}
			} else {
				best = dp(step+1, houses[step], nextNeighbour(nCount, prevColor, houses[step]))
			}
			if prevColor == 0 {
				return best
			}
			memo[step][prevColor-1][nCount] = best
		}
		return memo[step][prevColor-1][nCount]
	}
	res := dp(0, 0, 0)
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func nextNeighbour(count, prevColor, curColor int) int {
	if prevColor == curColor {
		return count
	}
	return count + 1
}
