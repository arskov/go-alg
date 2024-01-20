package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode122_1_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit_2_dp(prices)
	assert.Equal(t, 7, res)
}

func TestLeetcode122_1_2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfit_2_dp(prices)
	assert.Equal(t, 4, res)
}

func TestLeetcode122_1_3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit_2_dp(prices)
	assert.Equal(t, 0, res)
}

func maxProfit_2_dp(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	memo := make([][]int, 2)
	memo[0] = make([]int, n)
	memo[1] = make([]int, n)
	for i := 0; i < n; i++ {
		memo[0][i] = -1
		memo[1][i] = -1
	}

	var dp func(txActive int, i int) int
	dp = func(txActive int, i int) int {
		if i == n {
			return 0
		}
		if memo[txActive][i] == -1 {
			if txActive == 0 {
				a := -prices[i] + dp(txActive^1, i+1)
				b := dp(txActive, i+1)
				memo[txActive][i] = max(a, b)
			} else {
				a := prices[i] + dp(txActive^1, i+1)
				b := dp(txActive, i+1)
				memo[txActive][i] = max(a, b)
			}
		}
		return memo[txActive][i]
	}
	return dp(0, 0)
}
