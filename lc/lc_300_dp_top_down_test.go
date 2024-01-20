package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode300_1_0(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	assert.Equal(t, 4, res)
}

func TestLeetcode300_1_1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	res := lengthOfLIS(nums)
	assert.Equal(t, 4, res)
}

func TestLeetcode300_1_2(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	res := lengthOfLIS(nums)
	assert.Equal(t, 1, res)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dp func(i int) int
	dp = func(i int) int {
		if memo[i] == -1 {
			memo[i] = 1
			for j := i - 1; j >= 0; j-- {
				if nums[j] < nums[i] {
					memo[i] = max(memo[i], 1+dp(j))
				}
			}
		}
		return memo[i]
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = max(ans, dp(i))
	}
	return ans
}
