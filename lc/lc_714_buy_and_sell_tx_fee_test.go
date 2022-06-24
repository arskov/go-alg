package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode714_1_0(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	res := maxProfit_714(prices, fee)
	assert.Equal(t, 8, res)
}

func TestLeetcode714_1_1(t *testing.T) {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	res := maxProfit_714(prices, fee)
	assert.Equal(t, 6, res)
}

func maxProfit_714(prices []int, fee int) int {
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
	var dp func(txActive, step int) int
	dp = func(txActive, step int) int {
		if step == n {
			return 0
		}
		if memo[txActive][step] == -1 {
			if txActive == 1 {
				a := prices[step] - fee + dp(txActive^1, step+1)
				b := dp(txActive, step+1)
				memo[txActive][step] = mymath.Max(a, b)
			} else {
				a := -prices[step] + dp(txActive^1, step+1)
				b := dp(txActive, step+1)
				memo[txActive][step] = mymath.Max(a, b)
			}

		}
		return memo[txActive][step]
	}
	return dp(0, 0)
}
