package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode188_0_1(t *testing.T) {
	prices := []int{2, 4, 1}
	k := 2
	res := maxProfit_4_dp(k, prices)
	assert.Equal(t, 2, res)
}

func TestLeetcode188_0_2(t *testing.T) {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	res := maxProfit_4_dp(k, prices)
	assert.Equal(t, 7, res)
}

func maxProfit_4_dp(k int, prices []int) int {
	n := len(prices)
	if k == 0 || n == 0 {
		return 0
	}
	memo := make([][][]int, 2)
	memo[0] = make([][]int, k)
	memo[1] = make([][]int, k)
	for i := 0; i < k; i++ {
		memo[0][i] = make([]int, n)
		memo[1][i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[0][i][j] = -1
			memo[1][i][j] = -1
		}
	}
	var dp func(txActive, txCount, i int) int
	dp = func(txActive, txCount, i int) int {
		if i >= n || txCount >= k {
			return 0
		}
		if memo[txActive][txCount][i] == -1 {
			var a, b int
			if txActive == 0 {
				a = -prices[i] + dp(txActive^1, txCount, i+1)
				b = dp(txActive, txCount, i+1)
			} else {
				a = prices[i] + dp(txActive^1, txCount+1, i+1)
				b = dp(txActive, txCount, i+1)
			}
			memo[txActive][txCount][i] = mymath.Max(a, b)
		}
		return memo[txActive][txCount][i]
	}
	return dp(0, 0, 0)
}
