package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode518_1_0(t *testing.T) {
	coins := []int{1, 2, 5}
	res := change_518(5, coins)
	assert.Equal(t, 4, res)
}

func TestLeetcode518_1_1(t *testing.T) {
	coins := []int{2}
	res := change_518(3, coins)
	assert.Equal(t, 0, res)

}

func TestLeetcode518_1_2(t *testing.T) {
	coins := []int{10}
	res := change_518(10, coins)
	assert.Equal(t, 1, res)

}

func TestLeetcode518_1_3(t *testing.T) {
	coins := []int{3, 5, 7, 8, 9, 10, 11}
	res := change_518(500, coins)
	assert.Equal(t, 35_502_874, res)

}

func change_518(amount int, coins []int) int {
	memo := make([][]int, len(coins))
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dp func(int, int) int
	dp = func(amount, coinIdx int) int {
		if amount == 0 {
			return 1
		} else if amount < 0 {
			return 0
		}
		if memo[coinIdx][amount] == -1 {
			var count int
			for j := coinIdx; j >= 0; j-- {
				count += dp(amount-coins[j], j)
			}
			memo[coinIdx][amount] = count
		}
		return memo[coinIdx][amount]
	}
	return dp(amount, len(coins)-1)
}
