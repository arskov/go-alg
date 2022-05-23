package lc

import (
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode1143_1_0(t *testing.T) {
	text1 := "abc"
	text2 := "abc"
	res := longestCommonSubsequence(text1, text2)
	assert.Equal(t, 3, res)
}

func TestLeetcode1143_1_1(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	res := longestCommonSubsequence(text1, text2)
	assert.Equal(t, 3, res)
}

func TestLeetcode1143_1_2(t *testing.T) {
	text1 := "abc"
	text2 := "def"
	res := longestCommonSubsequence(text1, text2)
	assert.Equal(t, 0, res)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	n := len(text1)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if memo[i][j] == math.MaxInt32 {
			if text1[i] == text2[j] {
				memo[i][j] = 1 + dp(i+1, j+1)
			} else {
				memo[i][j] = mymath.Max(dp(i, j+1), dp(i+1, j))
			}
		}
		return memo[i][j]
	}
	return dp(0, 0)
}
