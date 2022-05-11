package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode787_2_0(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	res := findCheapestPrice_bellman_ford_dp(4, in, 0, 3, 1)
	assert.Equal(t, 700, res)
}

func TestLeetcode787_2_1(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_bellman_ford_dp(3, in, 0, 2, 1)
	assert.Equal(t, 200, res)
}

func TestLeetcode787_2_2(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_bellman_ford_dp(3, in, 0, 2, 0)
	assert.Equal(t, 500, res)
}

func TestLeetcode787_2_3(t *testing.T) {
	in := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{3, 7, 1},
		{0, 4, 3},
		{4, 5, 3},
		{5, 7, 3},
		{0, 6, 5},
		{6, 7, 100},
		{7, 8, 1},
	}
	res := findCheapestPrice_bellman_ford_dp(9, in, 0, 8, 3)
	assert.Equal(t, 10, res)
}

func TestLeetcode787_2_4(t *testing.T) {
	in := [][]int{
		{1, 2, 100},
		{3, 1, -150},
		{0, 1, 100},
		{2, 3, 100},
		{0, 3, 200},
		{0, 2, 500},
	}
	res := findCheapestPrice_bellman_ford_dp(4, in, 0, 1, 3)
	assert.Equal(t, 50, res)
}

func findCheapestPrice_bellman_ford_dp(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], math.MaxInt32)
		}
	}
	for i := 0; i < k+1; i++ {
		dp[i][src] = 0
		dp[i+1][src] = 0
		for _, edge := range flights {
			s := edge[0]
			d := edge[1]
			cost := edge[2]
			if dp[i][s]+cost < dp[i+1][d] {
				dp[i+1][d] = dp[i][s] + cost
			}
		}
	}
	if dp[k+1][dst] < math.MaxInt32 {
		return dp[k+1][dst]
	} else {
		return -1
	}
}
