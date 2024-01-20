package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1770_1_0(t *testing.T) {
	nums := []int{1, 2, 3}
	mult := []int{3, 2, 1}
	res := maximumScore(nums, mult)
	assert.Equal(t, 14, res)
}

func TestLeetcode1770_1_1(t *testing.T) {
	nums := []int{-5, -3, -3, -2, 7, 1}
	mult := []int{-10, -5, 3, 4, 6}
	res := maximumScore(nums, mult)
	assert.Equal(t, 102, res)
}

func maximumScore(nums []int, multipliers []int) int {
	mSize := len(multipliers)
	nSize := len(nums)
	memo := make([][]int, mSize)
	for i := 0; i < mSize; i++ {
		memo[i] = make([]int, mSize)
	}
	var dp func(multIdx, numStart int) int
	dp = func(multIdx, numStart int) int {
		if multIdx == mSize {
			return 0
		}
		if memo[multIdx][numStart] == 0 {
			numEnd := nSize - 1 - (multIdx - numStart)
			left := dp(multIdx+1, numStart)
			right := dp(multIdx+1, numStart+1)
			v := max(
				multipliers[multIdx]*nums[numEnd]+left,
				multipliers[multIdx]*nums[numStart]+right,
			)
			memo[multIdx][numStart] = v
		}
		return memo[multIdx][numStart]
	}
	return dp(0, 0)
}
