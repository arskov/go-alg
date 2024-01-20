package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode746_1_0(t *testing.T) {
	cost := []int{10, 15, 20}
	assert.Equal(t, 15, minCostClimbingStairs(cost))
}

func TestLeetcode746_1_1(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	assert.Equal(t, 6, minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	} else if n == 1 {
		return cost[0]
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
