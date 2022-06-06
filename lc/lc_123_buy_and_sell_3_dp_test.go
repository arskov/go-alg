package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode123_0_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit_3_dp(prices)
	assert.Equal(t, 7, res)
}

func TestLeetcode123_0_2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfit_3_dp(prices)
	assert.Equal(t, 4, res)
}

func TestLeetcode123_0_3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit_3_dp(prices)
	assert.Equal(t, 0, res)
}

func TestLeetcode123_0_4(t *testing.T) {
	prices := []int{2, 1, 4, 5, 2, 9, 7}
	res := maxProfit_3_dp(prices)
	assert.Equal(t, 11, res)
}

func maxProfit_3_dp(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	memo := make([][][]int, 2)
	memo[0] = make([][]int, 2)
	memo[1] = make([][]int, 2)
	memo[0][0] = make([]int, n)
	memo[0][1] = make([]int, n)
	memo[1][0] = make([]int, n)
	memo[1][1] = make([]int, n)

	for i := 0; i < n; i++ {
		memo[0][0][i] = -1
		memo[0][1][i] = -1
		memo[1][0][i] = -1
		memo[1][1][i] = -1
	}

	var dp func(txActive, txCount, i int) int
	dp = func(txActive, txCount, i int) int {
		if i >= n || txCount >= 2 {
			return 0
		}
		if memo[txActive][txCount][i] == -1 {
			if txActive == 0 {
				a := -prices[i] + dp(txActive^1, txCount, i+1)
				b := dp(txActive, txCount, i+1)
				memo[txActive][txCount][i] = mymath.Max(a, b)
			} else {
				a := prices[i] + dp(txActive^1, txCount+1, i+1)
				b := dp(txActive, txCount, i+1)
				memo[txActive][txCount][i] = mymath.Max(a, b)
			}
		}
		return memo[txActive][txCount][i]
	}
	return dp(0, 0, 0)
}
