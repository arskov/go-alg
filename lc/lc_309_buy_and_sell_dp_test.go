package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode309_0_1(t *testing.T) {
	prices := []int{1}
	res := maxProfit_309_dp(prices)
	assert.Equal(t, 0, res)
}

func TestLeetcode309_0_2(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	res := maxProfit_309_dp(prices)
	assert.Equal(t, 3, res)
}

func maxProfit_309_dp(prices []int) int {
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
	var dp func(txActive, i int) int
	dp = func(txActive, i int) int {
		if i >= n {
			return 0
		}
		if memo[txActive][i] == -1 {
			var a, b int
			if txActive == 1 {
				a = prices[i] + dp(txActive^1, i+2)
				b = dp(txActive, i+1)
			} else {
				a = -prices[i] + dp(txActive^1, i+1)
				b = dp(txActive, i+1)
			}
			memo[txActive][i] = mymath.Max(a, b)
		}
		return memo[txActive][i]
	}
	return dp(0, 0)
}
