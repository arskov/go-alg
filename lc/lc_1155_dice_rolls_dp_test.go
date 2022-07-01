package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1155_1_0(t *testing.T) {
	res := numRollsToTarget(1, 6, 3)
	assert.Equal(t, 1, res)
}

func TestLeetcode1155_1_1(t *testing.T) {
	res := numRollsToTarget(2, 6, 7)
	assert.Equal(t, 6, res)
}

func TestLeetcode1155_1_2(t *testing.T) {
	res := numRollsToTarget(30, 30, 500)
	assert.Equal(t, 222616187, res)
}

func numRollsToTarget(n int, k int, target int) int {
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(diceLeft, sum int) int
	dp = func(diceLeft, sum int) int {
		if sum == 0 {
			if diceLeft == 0 {
				return 1
			}
			return 0
		} else if sum < 0 || diceLeft < 0 {
			return 0
		}
		if memo[diceLeft][sum] == -1 {
			count := 0
			for i := 1; i <= k; i++ {
				count += dp(diceLeft-1, sum-i)
			}
			memo[diceLeft][sum] = count % 1_000_000_007
		}
		return memo[diceLeft][sum]
	}
	return dp(n, target)
}
