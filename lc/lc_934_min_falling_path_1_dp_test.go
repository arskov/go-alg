package lc

import (
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode934_1_0(t *testing.T) {
	mat := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	res := minFallingPathSum(mat)
	assert.Equal(t, 13, res)
}

func TestLeetcode934_1_1(t *testing.T) {
	mat := [][]int{
		{-19, 57},
		{-40, -5},
	}
	res := minFallingPathSum(mat)
	assert.Equal(t, -59, res)
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = matrix[i][j] + mymath.Min(dp[i-1][j], dp[i-1][j+1])
			} else if j == n-1 {
				dp[i][j] = matrix[i][j] + mymath.Min(dp[i-1][j], dp[i-1][j-1])
			} else {
				dp[i][j] = matrix[i][j] + mymath.Min(dp[i-1][j-1], mymath.Min(dp[i-1][j], dp[i-1][j+1]))
			}
		}
	}
	best := math.MaxInt32
	for i := 0; i < n; i++ {
		best = mymath.Min(best, dp[n-1][i])
	}
	return best
}
