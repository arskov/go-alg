package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const mod = 1_000_000_007

func TestLeetcode1220_1_0(t *testing.T) {
	res := countVowelPermutation(0)
	assert.Equal(t, 0, res)
}

func TestLeetcode1220_1_1(t *testing.T) {
	res := countVowelPermutation(1)
	assert.Equal(t, 5, res)
}

func TestLeetcode1220_1_2(t *testing.T) {
	res := countVowelPermutation(2)
	assert.Equal(t, 10, res)
}

func TestLeetcode1220_1_3(t *testing.T) {
	res := countVowelPermutation(3)
	assert.Equal(t, 19, res)
}

func countVowelPermutation(n int) int {
	if n <= 0 {
		return 0
	}
	memo := make([][]int, 5)
	for i := 0; i < 5; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	// 'a', 'e', 'i', 'o', 'u'
	//  0    1    2    3    4
	var dp func(prevLetter, l int) int
	dp = func(prevLetter, l int) int {
		if l == n-1 {
			return 1
		}
		if memo[prevLetter][l] == -1 {
			tmp := 0
			switch prevLetter {
			case 0:
				tmp = dp(1, l+1)
			case 1:
				tmp = dp(0, l+1) + dp(2, l+1)
			case 2:
				tmp = dp(0, l+1) + dp(1, l+1) + dp(3, l+1) + dp(4, l+1)
			case 3:
				tmp = dp(2, l+1) + dp(4, l+1)
			case 4:
				tmp = dp(0, l+1)
			}
			memo[prevLetter][l] = tmp % mod
		}
		return memo[prevLetter][l]
	}
	count := 0
	for i := 0; i < 5; i++ {
		count += dp(i, 0)
	}
	return count % mod
}
