package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode276_1_0(t *testing.T) {
	res := numWays_276(3, 2)
	assert.Equal(t, 6, res)
}

func TestLeetcode276_1_1(t *testing.T) {
	res := numWays_276(1, 1)
	assert.Equal(t, 1, res)
}

func TestLeetcode276_1_2(t *testing.T) {
	res := numWays_276(7, 2)
	assert.Equal(t, 42, res)
}

func numWays_276(n int, k int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dp func(int) int
	dp = func(step int) int {
		if step == 1 {
			return k
		} else if step == 2 {
			return k * k
		}
		if memo[step] == -1 {
			memo[step] = (k - 1) * (dp(step-1) + dp(step-2))
		}
		return memo[step]
	}
	return dp(n)
}
