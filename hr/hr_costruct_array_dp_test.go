package hr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHR_CountArray_1(t *testing.T) {
	res := countArray_1(4, 3, 2)
	assert.Equal(t, int64(3), res)
}

func TestHR_CountArray_2(t *testing.T) {
	res := countArray_2(4, 3, 2)
	assert.Equal(t, int64(3), res)
}

func TestHR_CountArray_3(t *testing.T) {
	res := countArray_3(4, 3, 2)
	assert.Equal(t, int64(3), res)
}

func countArray_1(n, k, x int32) int64 {
	dp := make([]int64, n+1)
	sp := make([]int64, n+1)
	dp[1] = 1
	var i int32
	for i = 2; i <= n; i++ {
		dp[i] = int64(k-1) * sp[i-1] % 1_000_000_007
		sp[i] = (dp[i-1] + sp[i-1]*int64(k-2)) % 1_000_000_007
	}
	if x == 1 {
		return dp[n]
	} else {
		return sp[n]
	}
}

func countArray_2(n, k, x int32) int64 {
	memo := make([][]int64, 2)
	memo[0] = make([]int64, n+1)
	memo[1] = make([]int64, n+1)
	var i int32
	for i = 0; i <= n; i++ {
		memo[0][i] = -1
		memo[1][i] = -1
	}
	var dp func(level, prev int32) int64
	dp = func(level, prev int32) int64 {
		if level == 2 {
			if prev == 0 {
				return 0
			} else {
				return 1
			}
		}
		if memo[prev][level] == -1 {
			var count int64
			if prev == 0 {
				count = int64(k-1) * dp(level-1, 1)
			} else {
				count = dp(level-1, 0) + int64(k-2)*dp(level-1, 1)
			}
			memo[prev][level] = count % 1_000_000_007
		}
		return memo[prev][level]
	}
	var prev int32
	if x == 1 {
		prev = 0
	} else {
		prev = 1
	}
	return dp(n, prev)
}

func countArray_3(n, k, x int32) int64 {
	// Return the number of ways to fill in the array.
	memo := make([][]int64, n+1)
	var i, j int32
	for i = 0; i < n+1; i++ {
		memo[i] = make([]int64, k+1)
		for j = 0; j < k+1; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(prev, i int32) int64
	dp = func(prev, i int32) int64 {
		if i == n {
			if prev == x {
				return 1
			} else {
				return 0
			}
		}
		if memo[i][prev] == -1 {
			var count int64 = 0
			var j int32
			for j = 1; j <= k; j++ {
				if j != prev {
					count += dp(j, i+1)
				}
			}
			memo[i][prev] = count % 1_000_000_007
		}
		return memo[i][prev]
		// return count % 1000000007
	}
	return dp(1, 1)
}
