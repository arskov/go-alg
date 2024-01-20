package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode221_1_0(t *testing.T) {
	matrix := [][]byte{
		{'0'},
	}
	res := maximalSquare(matrix)
	assert.Equal(t, 0, res)
}

func TestLeetcode221_1_1(t *testing.T) {
	matrix := [][]byte{
		{'1'},
	}
	res := maximalSquare(matrix)
	assert.Equal(t, 1, res)
}

func TestLeetcode221_1_2(t *testing.T) {
	matrix := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	res := maximalSquare(matrix)
	assert.Equal(t, 1, res)
}

func TestLeetcode221_1_3(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	res := maximalSquare(matrix)
	assert.Equal(t, 4, res)
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	maxSize := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSize {
					maxSize = dp[i][j]
				}
			}
		}
	}
	return maxSize * maxSize
}
