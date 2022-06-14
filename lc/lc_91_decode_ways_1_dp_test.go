package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode91_1_0(t *testing.T) {
	res := numDecodings("12")
	assert.Equal(t, 2, res)
}

func TestLeetcode91_1_1(t *testing.T) {
	res := numDecodings("226")
	assert.Equal(t, 3, res)
}

func TestLeetcode91_1_2(t *testing.T) {
	res := numDecodings("06")
	assert.Equal(t, 0, res)
}

func TestLeetcode91_1_3(t *testing.T) {
	res := numDecodings("1")
	assert.Equal(t, 1, res)
}

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dp func(int) int
	dp = func(i int) int {
		if i == n {
			return 1
		} else if i > n {
			return 0
		}
		if memo[i] == -1 {
			var count int
			val := s[i] - 48
			if val == 0 {
				count = 0
			} else if (val == 1) || (val == 2 && i+1 < n && s[i+1]-48 <= 6) {
				count = dp(i+1) + dp(i+2)
			} else {
				count = dp(i + 1)
			}
			memo[i] = count
		}
		return memo[i]
	}
	return dp(0)
}
