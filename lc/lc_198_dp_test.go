package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode198_1_0(t *testing.T) {
	res := rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, res)
}

func TestLeetcode198_1_1(t *testing.T) {
	res := rob([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, res)
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	dp[0] = nums[0]
	dp[1] = mymath.Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = mymath.Max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}
