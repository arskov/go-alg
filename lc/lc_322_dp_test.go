package lc

import (
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode322_1_0(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange(coins, amount)
	assert.Equal(t, 3, res)
}

func TestLeetcode322_1_1(t *testing.T) {
	coins := []int{2}
	amount := 3
	res := coinChange(coins, amount)
	assert.Equal(t, -1, res)
}

func TestLeetcode322_1_2(t *testing.T) {
	coins := []int{1}
	amount := 0
	res := coinChange(coins, amount)
	assert.Equal(t, 0, res)
}

func coinChange(coins []int, amount int) int {
	memo := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		memo[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	var dp func(i, am int) int
	dp = func(i, am int) int {
		if am == 0 {
			return 0
		}
		if am < 0 || i < 0 {
			return -1
		}
		if memo[i][am] == math.MaxInt32 {
			take := dp(i, am-coins[i])
			notTake := dp(i-1, am)
			result := -1
			if take == -1 && notTake == -1 {
				result = -1
			} else if take == -1 {
				result = notTake
			} else if notTake == -1 {
				result = take + 1
			} else {
				result = mymath.Min(take+1, notTake)
			}
			memo[i][am] = result
		}
		return memo[i][am]
	}
	return dp(len(coins)-1, amount)
}
